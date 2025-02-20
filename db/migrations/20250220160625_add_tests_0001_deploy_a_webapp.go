package migrations

func Init_20250220160625_add_tests_0001_deploy_a_webapp() Migration {
	return Migration{
		Id:        "20250220160625_add_tests_0001_deploy_a_webapp",
		Statement: "INSERT INTO tests(name, exercise_id, resource_type, resource_name, resource_attribute, resource_attribute_value) VALUES(?, ?, ?, ?, ?, ?)",
		Args: []any{
			"Nginx is installed",
			"0001-deploy-a-webapp",
			"Package",
			"nginx",
			"Status",
			"install ok installed",
		},
	}
}
