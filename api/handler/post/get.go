package post

import (
	"github.com/gofiber/fiber/v2"
)

func Get(ctx *fiber.Ctx) error {
	//TODO Implement Get Post
	/*
		response := ctx.Locals("response").(Response)
		searchID := ctx.Params("id", "")

		token := ctx.Locals("token").(auth.Token)
		myFirebaseID := token.UID
		if myFirebaseID == "" {
			response.Msg = MSG_FIREBASE_TOKEN
			response.Error = append(response.Error, ERROR_FIREBASE_TOKEN)
			response.send(fiber.StatusBadRequest)
			return nil
		}

		rows, err := getInfoFromRequest("SELECT todos.id,expire,creator,`check`,title,description,repetition,rotation,assignee,todos.`group` FROM todos INNER JOIN `groups` g on todos.`group` = g.id INNER JOIN members m on g.id = m.`group` INNER JOIN users u on m.user = u.id WHERE u.firebaseID = ? and todos.id = ?",
			myFirebaseID, searchID)
		if err != nil {
			response.Msg = MSG_DEFAULT
			response.Error = append(response.Error, "SELECT error")
			response.Error = append(response.Error, err.Error())
			response.send(fiber.StatusBadRequest)
			return nil
		}
		if !rows.Next() {
			response.Msg = "Es gibt keine Todo mit der ID auf die du zugriff hast"
			response.Error = append(response.Error, "rows empty")
			response.send(fiber.StatusBadRequest)
			return nil
		}
		todo, err := getTodoFromRow(&response, rows)
		if err != nil {
			response.send(fiber.StatusBadRequest)
			return nil
		}

		todo.Group, err = getGroupFromID(todo.Group.Id, myFirebaseID, &response)
		if err != nil {
			return nil
		}
		response.Content = todo
		response.send(200)
		return nil
		WITH RECURSIVE PostsTree AS (
			SELECT
		t.Id,
			t.CreatedAt,
			t.Creator,
			t.Group,
			t.Title,
			t.Content,
			t.Type,
			t.Parent,
			CAST(NULL AS CHAR(36)) AS RootId
		FROM
		posts t
		WHERE
		t.Parent IS NULL AND t.Group = ?
		UNION ALL
		SELECT
		t.Id,
			t.CreatedAt,
			t.Creator,
			t.Group,
			t.Title,
			t.Content,
			t.type,
		t.Parent,
			tt.Id AS RootId
		FROM
		posts t
		INNER JOIN
		PostsTree tt ON t.Parent = tt.Id
		)
		SELECT
		RootId,
			JSON_ARRAYAGG(JSON_OBJECT(
				'Id', Id,
				'CreatedAt', CreatedAt,
				'Creator', Creator,
				'Group', `Group`,
				'Title', Title,
				'Content', Content,
				'Type', type,
		'Parent', Parent
		)) AS Posts
		FROM
		PostsTree
		GROUP BY
		RootId;


	*/
	return nil
}
