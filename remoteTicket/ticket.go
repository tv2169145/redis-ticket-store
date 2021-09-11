package remoteTicket

import "github.com/gomodule/redigo/redis"

const LuaScript = `
        local ticket_key = KEYS[1]
        local ticket_total_key = ARGV[1]
        local ticket_sold_key = ARGV[2]
        local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
        local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
		-- 查看是否还有余票,增加订单数量,返回结果值
        if(ticket_total_nums >= ticket_sold_nums) then
            return redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
        end
        return 0
`

// redis訂單Key
type RemoteTicketsKeys struct {
	HashKey string	            // redis中hash結構的key
	TicketTotalKey string	    // hash結構中總庫存量的key
	QuantityOfOrderKey string	// hash結構中已有訂單數量的key (賣出量)
}

func NewRemoteTicketKeys() *RemoteTicketsKeys {
	return &RemoteTicketsKeys{
		HashKey: "ticket_hash_key",
		TicketTotalKey: "ticket_total_nums",
		QuantityOfOrderKey: "ticket_sold_nums",
	}
}

// redis統一扣庫存
func (hashKeys *RemoteTicketsKeys) RemoteDeductionTicket(conn redis.Conn) bool {
	lua := redis.NewScript(1, LuaScript)
	result, err := redis.Int(lua.Do(conn, hashKeys.HashKey, hashKeys.TicketTotalKey, hashKeys.QuantityOfOrderKey))
	if err != nil {
		return false
	}
	return result != 0
}
