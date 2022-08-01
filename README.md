## Validation check between Frontend and Backend 

변수들의 유효성 검사는 프론트엔드에서 진행한다.
백엔드는 모든 변수들이 정확히 들어온다고 가정한다.

## Database

### OAuth를 사용하지 않은 로그인 사용자 정보 테이블 구조

| key      | description                   |
| -------- | ----------------------------- |
| id       | Primary key                   |
| user_id  | User identifiable information |
| name     | Duplicable information        |
| gender   | Gender                        |
| password | Hashed by SHA256              |


### 네이버 로그인 사용자 정보 테이블 구조

``` sql
CREATE TABLE `SNS_INFO` (
  `id` int(11) NOT NULL,
  `sns_id` varchar(255) NOT NULL,
  `sns_type` varchar(10)  NULL,
  `sns_name` varchar(255)  NULL,
  `sns_profile` varchar(255)  NULL,
  `sns_connect_date` datetime  NULL,
  KEY `idx01_id` (`id`),
  KEY `idx02_sns_id` (`sns_id`),
  CONSTRAINT `id` FOREIGN KEY (`id`) REFERENCES `USERS` (`id`)
);
```