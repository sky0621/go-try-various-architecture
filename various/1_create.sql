# システムにログインして商品を注文するユーザーの属性
CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(64) NOT NULL,
  `name` varchar(256) NOT NULL,
  `mail` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

# 商品をユニークに特定しうる属性
CREATE TABLE IF NOT EXISTS `item` (
  `id` varchar(64) NOT NULL,
  `name` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

#
CREATE TABLE IF NOT EXISTS `item_status` (
  `item_id` varchar(64) NOT NULL,
  `name` varchar(256) NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

# 任意のユーザによる商品群の注文
CREATE TABLE IF NOT EXISTS `order` (
  `id` varchar(64) NOT NULL,
  `user_id` varchar(64) NOT NULL,
  `nums` int NOT NULL,
  `total_price` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `order_item` (
  `id` varchar(64) NOT NULL,
  `order_id` varchar(64) NOT NULL,
  `item_id` varchar(64) NOT NULL,
  `nums` int NOT NULL,
  `total_price` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;
