package ui

import (
	"context"
	"database/sql"
	"os"
	"reflect"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/living-etc/sightseer.go/clients/ssh"
)

type Test struct {
	Id                     int
	Name                   string
	ExerciseId             string
	ResourceType           string
	ResourceName           string
	ResourceAttribute      string
	ResourceAttributeValue string
	Negation               bool
	Result                 bool
}

type TestResult struct {
	Value bool
}

func (t Test) Run(logger *log.Logger) (TestResult, error) {
	logger.Info("[Running test]", "Name", t.Name)

	privateKey, err := os.ReadFile("/Users/chris/.ssh/id_ed25519")
	if err != nil {
		logger.Debug(err)
	}

	client, err := ssh.NewSshClient(
		privateKey,
		"192.168.86.96",
		"22",
		"chris",
		"ubuntu2404",
	)
	if err != nil {
		logger.Error(err)
	}

	resource := reflect.ValueOf(client).
		MethodByName(t.ResourceType).
		Call([]reflect.Value{reflect.ValueOf(t.ResourceName)})[0].Interface()

	attribute := reflect.Indirect(reflect.ValueOf(resource)).
		FieldByName(t.ResourceAttribute).
		Interface().(string)

	res := strings.Compare(attribute, t.ResourceAttributeValue)

	var result bool
	if res == 0 {
		result = true
	}

	return TestResult{
		Value: result,
	}, nil
}

func (t Test) RecordResult(conn *sql.DB, result TestResult) error {
	ctx := context.Background()

	_, err := conn.QueryContext(
		ctx,
		"UPDATE tests SET result = ? WHERE id = ?",
		result.Value,
		t.Id,
	)

	return err
}
