package pc

const ErrorOk = 1001 // success

const ErrorParam = 2000 // 参数有误

const ErrorSystem = 10000 // 系统异常
const ErrorJsonMarshal = 10001 // json化异常

const ErrorRedisGet = 10101 // redis get有误
const ErrorRedisSet = 10102 // redis set有误
const ErrorRedisDel = 10103 // redis del有误
const ErrorRedisSetEx = 10104 // redis setEx有误
const ErrorRedisDecr = 10105 // redis decr有误
const ErrorRedisIncr = 10106 // redis incr有误

const ErrorMysqlSelect = 10201 // mysql select有误
const ErrorMysqlUpdate = 10202 // mysql update有误
const ErrorMysqlCreate = 10203 // mysql create有误