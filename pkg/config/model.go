package config

/**
 * ConfigurationLoader contains all methods to load/save configuration files
 */
type ConfigurationLoader struct {
}

/**
 * The Project Configuration
 */
type ProjectConfigrationFile struct {
	Images []struct {
		// name of the container
		Name string

		// description for the  container
		Description string

		// the commands provided by the image
		Provides []string

		// container image
		Image string

		// target directory to mount your project inside of the container
		Directory string `default:"/project"`

		// wrap the executed command inside of the container into a shell (ex. if you use globs)
		Shell string `default:"none"`

		// commands that should run in the container before the actual command is executed
		BeforeScript []string `yaml:"before_script"`

		// Caching of container-directories
		Caching []CachingEntry `yaml:"cache"`

		// the command scope (internal use only) - global or project
		Scope string
	}
}

type CachingEntry struct {

	/**
	 * Name of the caching entry
	 */
	Name string `yaml:"name",default:""`

	/**
	 * Directory inside of the container that should be mounted on the host within the cache directory
	 */
	ContainerDirectory string `yaml:"directory",default:""`
}

/**
 * The EnvCLI Configuration
 */
type PropertyConfigurationFile struct {
	Properties map[string]string
}