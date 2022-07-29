## Validation check between Frontend and Backend 

변수들의 유효성 검사는 프론트엔드에서 진행한다.
백엔드는 모든 변수들이 정확히 들어온다고 가정한다.

## Database

- User : Sqlite3

### User

| key      | description                   |
| -------- | ----------------------------- |
| id       | Primary key                   |
| user_id  | User identifiable information |
| name     | Duplicable information        |
| gender   | Gender                        |
| password | Hashed by SHA256              |