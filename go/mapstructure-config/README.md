Variable Configuration with Golang `mapstructure`
=================================================

Goal:  prototype a mechanism where a large YAML file representing a
configuration file can be loaded and then subparts of that file are used to
populate golang struct's that represent component configuration.

Note:  this prototype was constructed not long after mapstructure migrated from
github.com/mitchellh/mapstructure to github.com/go-viper/mapstructure.  At this
time, the go-viper/mapstructure is the blessed successor.
