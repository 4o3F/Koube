CREATE TABLE `audience` (
    `aid` int UNIQUE PRIMARY KEY,
    `entrance_code` varchar(255) NOT NULL,
    `verify_code` varchar(255) NOT NULL UNIQUE,
    `entered` boolean NOT NULL DEFAULT false
)