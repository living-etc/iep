INSERT INTO exercises(
  id,
  name,
  description
) VALUES(
  "0001-deploy-a-webapp",
  "Deploy a Web Server with Nginx and AWS",
  "Learn how to put a website on the internet using Nginx and run it on an EC2 instance."
);

INSERT INTO tests(
  name,
  exercise_id,
  resource_type,
  resource_name,
  resource_attribute,
  resource_attribute_value,
  negation
) VALUES(
  "Nginx is installed",
  "0001-deploy-a-webapp",
  "Package",
  "nginx",
  "Status",
  "install ok installed",
  1
);

INSERT INTO tests(
  name,
  exercise_id,
  resource_type,
  resource_name,
  resource_attribute,
  resource_attribute_value,
  negation
) VALUES(
	"Nginx service is running",
	"0001-deploy-a-webapp",
	"Service",
	"nginx",
	"ActiveState",
	"active",
	1
);
