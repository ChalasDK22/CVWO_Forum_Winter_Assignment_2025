package topic

import (
	"context"
	"fmt"

	//"database/sql"
	//"fmt"
	"testing"

	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/internal/testutil"
	"chalas.com/forum_project/API/pkg/internalsql"
)

func TestTopicRepository_InsertTopic(t *testing.T) {
	chalasconfig, err := config.ConfigLoad()
	chalas_mydb, err := internalsql.ConnectAPI_MYSQL(chalasconfig)
	if err != nil {
		fmt.Println(err)
	}
	test_repo := NewTopicRepository(chalas_mydb)
	test_topic := testutil.NewTopicTestingModel()

	id, err := test_repo.InsertTopic(context.Background(), test_topic)
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatal("Topic ID should be other than 0")
	}
}

func TestTopicRepository_UpdateTopic(t *testing.T) {
	chalasconfig, err := config.ConfigLoad()
	chalasMyDB, err := internalsql.ConnectAPI_MYSQL(chalasconfig)
	if err != nil {
		fmt.Println(err)
	}
	testRepo := NewTopicRepository(chalasMyDB)
	testTopic := testutil.NewUpdatedTopicTestModel()

	err = testRepo.UpdateRepoTopic(context.Background(), 16, testTopic)
	if err != nil {
		t.Fatal(err)
	}

}
