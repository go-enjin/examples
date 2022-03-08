CREATE TABLE `content` (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `url` text,
  `slug` text,
  `title` text,
  `format` text,
  `summary` text,
  `description` text,
  `archetype` text,
  `section` text,
  `content` text,
  `language` text,
  `front_matter` text,
  `path` text,
  `layout` text,
  PRIMARY KEY (`id`)
);
CREATE INDEX `idx_content_title` ON `content`(`title`);
CREATE INDEX `idx_content_url` ON `content`(`url`);
CREATE INDEX `idx_content_deleted_at` ON `content`(`deleted_at`);
