
CREATE TABLE users (
                       id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID，主键，自增',
                       username VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户名',
                       phone VARCHAR(50) NOT NULL DEFAULT '' COMMENT '手机号',
                       password VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
                       salt  varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
                       role ENUM('super_admin', 'manager', 'operator') NOT NULL DEFAULT 'operator' COMMENT '用户角色：super_admin, manager, operator',
                       rate decimal(10,4) NOT NULL default 0 COMMENT '费用比例',
                       parent_id BIGINT Default 0 COMMENT '上级用户ID，指向所属工作室或超级管理员',
                       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                       UNIQUE KEY phone (phone)
) COMMENT='用户表';


CREATE TABLE channels (
                          id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '频道ID，主键，自增',
                          uid BIGINT NOT NULL DEFAULT 0 comment '用户ID',
                          account_name VARCHAR(100) NOT NULL COMMENT '账号名称（YouTube登录的Gmail）',
                          channel_name VARCHAR(100) NOT NULL COMMENT '频道名称',
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          last_sync_time DATETIME COMMENT '最后一次同步时间',
                          is_synced_today ENUM('yes', 'no') NOT NULL DEFAULT 'no' COMMENT '当日是否已同步视频数据'
) COMMENT='频道表';


CREATE TABLE videos (
                        id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '视频ID，主键，自增',
                        channel_id BIGINT NOT NULL COMMENT '所属频道ID',
                        video_cover VARCHAR(255) NOT NULL COMMENT '视频封面URL',
                        video_title VARCHAR(255) NOT NULL COMMENT '视频标题',
                        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        song_name VARCHAR(255) COMMENT '歌曲名（视频挂载的BGM）',
                        has_copyright_issue ENUM('yes', 'no') NOT NULL DEFAULT 'no' COMMENT '是否有版权问题',
                        is_matched_to_library ENUM('yes', 'no') NOT NULL DEFAULT 'no' COMMENT '是否匹配曲库',
                        isrc varchar(50) not null default ''
) COMMENT='视频表';



CREATE TABLE video_stats (
                             id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '统计ID，主键，自增',
                             video_id BIGINT NOT NULL COMMENT '视频ID',
                             total_views BIGINT NOT NULL DEFAULT 0 COMMENT '总播放量',
                             premium_views BIGINT NOT NULL DEFAULT 0 COMMENT 'Premium播放量',
                             likes BIGINT NOT NULL DEFAULT 0 COMMENT '点赞量',
                             comments BIGINT NOT NULL DEFAULT 0 COMMENT '评论数',
                             estimated_revenue DECIMAL(10,2) NOT NULL DEFAULT 0.00 COMMENT '预估收益',
                             country varchar(50) NOT NULL DEFAULT '' comment '国家',
                             stat_date varchar(50) NOT NULL COMMENT '统计日期'
) COMMENT='视频统计数据表';

CREATE TABLE channel_stats (
                               id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '统计ID，主键，自增',
                               channel_id BIGINT NOT NULL COMMENT '频道ID',
                               total_views BIGINT NOT NULL DEFAULT 0 COMMENT '总播放量',
                               premium_views BIGINT NOT NULL DEFAULT 0 COMMENT 'Premium播放量',
                               estimated_revenue DECIMAL(10,2) NOT NULL DEFAULT 0.00 COMMENT '预估收益',
                               country varchar(50) NOT NULL DEFAULT '' comment '国家',
                               stat_date varchar(50) NOT NULL COMMENT '统计日期'
) COMMENT='频道统计数据表';

CREATE TABLE country_rate (
                              id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '统计ID，主键，自增',
                              country_rate DECIMAL(10,4) NOT NULL DEFAULT 0.00 COMMENT '费率',
                              country varchar(50) NOT NULL DEFAULT '' comment '国家',
                              updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='国家费率表';


CREATE TABLE operation_logs (
                                id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '日志ID，主键，自增',
                                uid BIGINT NOT NULL DEFAULT 0 COMMENT '操作用户ID',
                                action int NOT NULL default 0 COMMENT '操作类型',
                                details TEXT NOT NULL COMMENT '操作详情',
                                created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间'
) COMMENT='操作日志表';

CREATE TABLE music (
                       id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '音乐ID，主键，自增',
                       uid BIGINT NOT NULL DEFAULT 0 comment '用户ID',
                       cover_url VARCHAR(255) NOT NULL COMMENT '歌曲封面URL',
                       song_name VARCHAR(255) NOT NULL COMMENT '歌曲名',
                       artist_name VARCHAR(255) NOT NULL COMMENT '艺人名',
                       isrc VARCHAR(50) NOT NULL COMMENT 'ISRC（国际标准录音代码）',
                       music_link VARCHAR(255) NOT NULL COMMENT '音乐链接',
                       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='音乐表';


CREATE TABLE tags (
                      id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '标签ID，主键，自增',
                      tag_name VARCHAR(100) NOT NULL COMMENT '标签名称',
                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                      UNIQUE KEY unique_tag_name (tag_name)
) COMMENT='标签表';

CREATE TABLE music_tags (
                            id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '关联ID，主键，自增',
                            music_id BIGINT NOT NULL COMMENT '音乐ID',
                            tag_id BIGINT NOT NULL COMMENT '标签ID',
                            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            FOREIGN KEY (music_id) REFERENCES music(id),
                            FOREIGN KEY (tag_id) REFERENCES tags(id),
                            UNIQUE KEY unique_music_tag (music_id, tag_id)
) COMMENT='音乐标签关联表';



