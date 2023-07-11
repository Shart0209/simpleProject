INSERT INTO suppliers
VALUES (DEFAULT, 'ПАО МТС', 'Sponge Bob', 'Чита', 'Забайкальский край'),
       (DEFAULT, 'ПАО РОСТЕЛЕКОМ', 'Sponge Bob', 'Чита', 'Забайкальский край'),
       (DEFAULT, 'ПАО МЕГАФОН', 'Sponge Bob', 'Чита', 'Забайкальский край');

INSERT INTO authors (login, name, role, pswd_hash, email)
VALUES ('test', 'test', 'user', '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08','example@gmail.com');


WITH new_contract as  (
	INSERT INTO contracts (title, numb, price, date, start_date, end_date, description, files)
	VALUES (
		'Оказание услуг Интернет','ЧИТ-13-45', '152000', '2023-07-01', '2023-07-01','2023-12-31', 'test test description', '[
  {
    "id": "e64b7418-1478-4e78-a866-74e0c4b77419-0.txt",
    "name": "гос контракт тест.txt",
    "path": "/upload/ГК № ЧИТ-13-45",
    "size": 30936
  }
]'
	)
	RETURNING contract_id 
	)
	INSERT INTO commons
	SELECT contract_id, 2, 17, 1, 1, true, '2023-07-01', '2023-07-01' from new_contract
