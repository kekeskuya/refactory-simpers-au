package config

import "refactory-simpers-au/docs"

func InitSwagger(env *EnvironmentVariable) {
	docs.SwaggerInfo.Title = env.Swagger.Title
	docs.SwaggerInfo.Description = env.Swagger.Description
	docs.SwaggerInfo.Version = env.Swagger.Version
	docs.SwaggerInfo.Host = env.Swagger.Host
}
