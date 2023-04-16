package main

import (
	"io/fs"
	"io/ioutil"

	config "migrate.com/src/config"
)

func main() {

	db := config.DB()

	checkfile, err := ioutil.ReadFile("sudahmigrasi.txt")

	if err != nil || string(checkfile) != "" {
		db.Exec("CREATE TABLE `activities` (`id` int(11) NOT NULL, `title` varchar(100) NOT NULL, `email` varchar(100) NOT NULL, `created_at` timestamp NOT NULL DEFAULT current_timestamp(), `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;");
		db.Exec("INSERT INTO `activities` (`id`, `title`, `email`, `created_at`, `updated_at`) VALUES (29, '', 'sdfdsfdsf', '2023-04-15 14:41:21', '2023-04-15 14:41:21');");
		db.Exec("CREATE TABLE `todos` (`id` int(11) NOT NULL,`title` varchar(100) NOT NULL,`activity_group_id` int(11) NOT NULL,`is_active` tinyint(1) NOT NULL DEFAULT 0,`priority` enum('','very-high','medium','low') DEFAULT 'very-high',`created_at` timestamp NOT NULL DEFAULT current_timestamp(),`updated_at` timestamp NOT NULL DEFAULT current_timestamp()) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;");
		db.Exec("INSERT INTO `todos` (`id`, `title`, `activity_group_id`, `is_active`, `priority`, `created_at`, `updated_at`) VALUES (1000000020, '', 29, 0, 'very-high', '2023-04-15 14:41:53', '2023-04-15 14:41:53');");
		db.Exec("ALTER TABLE `activities`ADD PRIMARY KEY (`id`),ADD UNIQUE KEY `email` (`email`);");
		db.Exec("ALTER TABLE `todos`ADD PRIMARY KEY (`id`),ADD KEY `activity_group_id` (`activity_group_id`);");
		db.Exec("ALTER TABLE `activities`MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;");
		db.Exec("ALTER TABLE `todos`MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1000000032;");
		db.Exec("ALTER TABLE `todos`ADD CONSTRAINT `todos_ibfk_1` FOREIGN KEY (`activity_group_id`) REFERENCES `activities` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;");

		ioutil.WriteFile("sudahmigrasi.txt", []byte("is string"), fs.ModeAppend)
	}
}
