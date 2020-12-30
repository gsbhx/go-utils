package redis

import redigo "github.com/garyburd/redigo/redis"

type Cbk struct {
	Cback interface{}
	Err   error
}

func (c *Cbk) String() (string, error) {
	return redigo.String(c.Cback, c.Err)
}
func (c *Cbk) Int() (int, error) {
	return redigo.Int(c.Cback, c.Err)
}
func (c *Cbk) Bool() (bool, error) {
	return redigo.Bool(c.Cback, c.Err)
}
func (c *Cbk) Slice() (slices []string, err error) {
	res, err := redigo.ByteSlices(c.Cback, c.Err)
	if err != nil {
		return nil, err
	}
	for _, v := range res {

		slices = append(slices, string(v))
	}
	return
}
func (c *Cbk) ScanSlice() (iter int, valMap map[string]string, err error) {
	cbk := c.Cback.([]interface{})
	iter, _ = redigo.Int(cbk[0], nil)
	valMap = make(map[string]string)
	cbkVal := cbk[1].([]interface{})
	for i := 0; i < len(cbkVal); i++ {
		if i%2 > 0 {
			valMap[string(cbkVal[i-1].([]byte))] = string(cbkVal[i].([]byte))
		}
	}
	return
}
func (c *Cbk) Error() error {
	return c.Err
}
func GetCbk(cbk interface{}, err error) *Cbk {
	return &Cbk{
		Cback: cbk,
		Err:   err,
	}
}
