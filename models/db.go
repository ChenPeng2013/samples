package models

import (
    "gopkg.in/mgo.v2"
    "github.com/spf13/viper"
    "fmt"
    "reflect"
)

type DB struct {
    Sess *mgo.Session
}

var (
    gDB *DB
)

func getDB() *DB {
    return gDB
}

func init() {
    viper.SetEnvPrefix("Samples")
    viper.AutomaticEnv()
    IPAddr := viper.GetString("DBHOST")
    fmt.Println(viper.GetString("DBHOST"))
    IPPort := viper.GetString("DBPORT")
    fmt.Println(viper.GetString("DBPORT"))

    gSess, err := mgo.Dial(IPAddr + ":" + IPPort)
    if err != nil {
        fmt.Println(err)
    }
    gDB = &DB {
        Sess:  gSess,
    }
}

func (db *DB) FindAll(slice interface{}) error {
    sliceData := reflect.Indirect(reflect.ValueOf(slice))
    //fmt.Printf("%+v", element)
    if sliceData.Kind() != reflect.Slice {
        return fmt.Errorf("请输入数组的指针对象")
    }

    col := db.Sess.DB("Test").C("Task")
    query := col.Find(nil)

    if err := query.All(slice); err != nil {
        return err
    }
    return nil
}

func (db *DB) Insert(docs interface{}) error {
    collection := db.Sess.DB("Test").C("Task")
    if err := collection.Insert(docs); err != nil {
        fmt.Println(err)
        return  err
    }
    return nil
}

func (db *DB) UpdateOne(sel interface{}, up interface{}) error {
    database := db.Sess.DB("Test")
    collection := database.C("Task")

    return collection.Update(sel, up)
}