package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sync"
)

var (
	checker *Checker
)

const (
	defaultConf = "D:\\go\\src\\checkgo\\conf\\gos.json"
)

func init() {
	checker = &Checker{
		lock: &sync.Mutex{},
	}
	checker.LoadConf(defaultConf)
}

type Checker struct {
	lock     *sync.Mutex
	jsonConf Conf
}

//GetChecker 获取校验器
func GetChecker() *Checker {
	return checker
}

//LoadConf 获取校验器
//Conf conf 配置加载
func (c *Checker) LoadConf(file string) error {

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	var conf Conf
	if err = json.Unmarshal(bytes, &conf); err != nil {
		return err
	}

	log.Println(string(bytes))

	if err := conf.check(); err != nil {
		return err
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	c.jsonConf = conf
	return nil
}

//GetActivity 获取活动配配置
//string activityCode 配置加载
func (c *Checker) GetActivity(activityCode string) (*Activity, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.jsonConf.Activities != nil && c.jsonConf.Activities[activityCode] != nil {
		return c.jsonConf.Activities[activityCode], nil
	}
	return nil, errors.New("activity not found")
}

//GetAllActivity 获取活动配配置
//string activityCode 配置加载
func (c *Checker) GetAllActivity() ([]*Activity, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	act := make([]*Activity, 0, len(c.jsonConf.Activities))
	if c.jsonConf.Activities != nil {
		for _, activity := range c.jsonConf.Activities {
			act = append(act, activity)
		}
	}
	return act, nil
}
