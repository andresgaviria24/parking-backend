CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `userName` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `parking` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `fechaIngreso` datetime DEFAULT NULL,
  `placa` varchar(100) DEFAULT NULL,
  `valor` float DEFAULT NULL,
  `nombrePropietario` varchar(100) DEFAULT NULL,
  `fechaRegistro` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci