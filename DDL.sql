CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `stock` int NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `ticket` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `id_product` int NOT NULL,
  PRIMARY KEY (`id`)
);
