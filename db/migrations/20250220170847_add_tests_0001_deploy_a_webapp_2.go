package migrations

func Init_20250220170847_add_tests_0001_deploy_a_webapp_2() Migration {
	return Migration{
		Id:        "20250220170847_add_tests_0001_deploy_a_webapp_2",
		Statement: "INSERT INTO tests(name, exercise_id, resource_type, resource_name, resource_attribute, resource_attribute_value, negation) VALUES(?, ?, ?, ?, ?, ?, ?)",
		Args: []any{
			"Nginx service is running",
			"0001-deploy-a-webapp",
			"Service",
			"nginx",
			"ActiveState",
			"active",
			1,
		},
	}
}
