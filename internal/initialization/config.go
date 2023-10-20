package initialization

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"

    "github.com/spf13/viper"
)

type Config struct {
    HttpPort     int    `json:"httpPort"`
    DbHost       string `json:"dbHost"`
    DbPort       int    `json:"dbPort"`
    DbDatabase   string `json:"dbDatabase"`
    DbUsername   string `json:"dbUsername"`
    DbPassword   string `json:"dbPassword"`
}

var AppConfig Config

func LoadConfig(cfg string) Config {
    viper.SetConfigFile(cfg)
    viper.ReadInConfig()
    viper.AutomaticEnv()
    content, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        fmt.Println("Error reading file:", err)
    }
    fmt.Println(string(content))
    AppConfig = Config{
        HttpPort:   getViperIntValue("HTTP_PORT", 9000),
        DbHost:     getViperStringValue("DB_HOST", "localhost"),
        DbPort:     getViperIntValue("DB_PORT", 3306),
        DbDatabase: getViperStringValue("DB_DATABASE", ""),
        DbUsername: getViperStringValue("DB_USERNAME", ""),
        DbPassword: getViperStringValue("DB_PASSWORD", ""), // 使用默认值或隐藏值
    }
    fmt.Println("读取到的配置信息：", AppConfig)

    return AppConfig
}

func getViperStringValue(key string, defaultValue string) string {
    value := viper.GetString(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getViperStringArray(key string, defaultValue []string) []string {
    value := viper.GetString(key)
    if value == "" {
        return defaultValue
    }
    raw := strings.Split(value, ",")
    return raw
}

func getViperIntValue(key string, defaultValue int) int {
    value := viper.GetString(key)
    if value == "" {
        return defaultValue
    }
    intValue, err := strconv.Atoi(value)
    if err != nil {
        fmt.Printf("Invalid value for %s, using default value %d\n", key, defaultValue)
        return defaultValue
    }
    return intValue
}

func getViperBoolValue(key string, defaultValue bool) bool {
    value := viper.GetString(key)
    if value == "" {
        return defaultValue
    }
    boolValue, err := strconv.ParseBool(value)
    if err != nil {
        fmt.Printf("Invalid value for %s, using default value %v\n", key, defaultValue)
        return defaultValue
    }
    return boolValue
}
