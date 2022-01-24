package config

import (
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var DefaultPath = "./conf/conf.yaml"

func InitConfig(p string) {
	viper.SetConfigFile(p)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: err=%v", err)
	}
}

func WatchConfig() {
	viper.WatchConfig()
}

func AllowEmptyEnv(allowEmptyEnv bool) {
	viper.AllowEmptyEnv(allowEmptyEnv)
}

func AddConfigPath(in string) {
	viper.AddConfigPath(in)
}

func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error {
	return viper.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
}

func SetTypeByDefaultValue(enable bool) {
	viper.SetTypeByDefaultValue(enable)
}

func GetViper() *viper.Viper {
	return viper.GetViper()
}

func Get(key string) interface{} { return viper.Get(key) }

func Sub(key string) *viper.Viper { return viper.Sub(key) }

func GetString(key string) string { return viper.GetString(key) }

func GetBool(key string) bool { return viper.GetBool(key) }

func GetInt(key string) int { return viper.GetInt(key) }

func GetInt32(key string) int32 { return viper.GetInt32(key) }

func GetInt64(key string) int64 { return viper.GetInt64(key) }

func GetUint(key string) uint { return viper.GetUint(key) }

func GetUint32(key string) uint32 { return viper.GetUint32(key) }

func GetUint64(key string) uint64 { return viper.GetUint64(key) }

func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

func GetTime(key string) time.Time { return viper.GetTime(key) }

func GetDuration(key string) time.Duration { return viper.GetDuration(key) }

func GetIntSlice(key string) []int { return viper.GetIntSlice(key) }

func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

func GetStringMap(key string) map[string]interface{} { return viper.GetStringMap(key) }

func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }

func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

func GetSizeInBytes(key string) uint { return viper.GetSizeInBytes(key) }

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.UnmarshalKey(key, rawVal, opts...)
}

func Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.Unmarshal(rawVal, opts...)
}

func UnmarshalExact(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.UnmarshalExact(rawVal, opts...)
}

func BindPFlags(flags *pflag.FlagSet) error { return viper.BindPFlags(flags) }

func BindPFlag(key string, flag *pflag.Flag) error { return viper.BindPFlag(key, flag) }

func BindFlagValues(flags viper.FlagValueSet) error { return viper.BindFlagValues(flags) }

func BindFlagValue(key string, flag viper.FlagValue) error { return viper.BindFlagValue(key, flag) }

func BindEnv(input ...string) error { return viper.BindEnv(input...) }

func IsSet(key string) bool { return viper.IsSet(key) }

func AutomaticEnv() { viper.AutomaticEnv() }

func SetEnvKeyReplacer(r *strings.Replacer) { viper.SetEnvKeyReplacer(r) }

func RegisterAlias(alias string, key string) { viper.RegisterAlias(alias, key) }

func InConfig(key string) bool { return viper.InConfig(key) }

func SetDefault(key string, value interface{}) { viper.SetDefault(key, value) }

func Set(key string, value interface{}) { viper.Set(key, value) }

func MergeInConfig() error { return viper.MergeInConfig() }

func ReadConfig(in io.Reader) error { return viper.ReadConfig(in) }

func MergeConfig(in io.Reader) error { return viper.MergeConfig(in) }

func MergeConfigMap(cfg map[string]interface{}) error { return viper.MergeConfigMap(cfg) }

func WriteConfig() error { return viper.WriteConfig() }

func SafeWriteConfig() error { return viper.SafeWriteConfig() }

func WriteConfigAs(filename string) error { return viper.WriteConfigAs(filename) }

func SafeWriteConfigAs(filename string) error { return viper.SafeWriteConfigAs(filename) }

func ReadRemoteConfig() error { return viper.ReadRemoteConfig() }

func WatchRemoteConfig() error { return viper.WatchRemoteConfig() }

func AllKeys() []string { return viper.AllKeys() }

func AllSettings() map[string]interface{} { return viper.AllSettings() }

func SetFs(fs afero.Fs) { viper.SetFs(fs) }

func SetConfigName(in string) { viper.SetConfigName(in) }

func SetConfigType(in string) { viper.SetConfigType(in) }

func SetConfigPermissions(perm os.FileMode) { viper.SetConfigPermissions(perm) }

func IniLoadOptions(in ini.LoadOptions) viper.Option {
	return viper.IniLoadOptions(in)
}

func Debug() { viper.Debug() }
