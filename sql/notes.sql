CREATE TABLE notes (
                       id INT AUTO_INCREMENT PRIMARY KEY,      -- 笔记 ID
                       user_id INT NOT NULL,                   -- 用户 ID (多用户支持)
                       title VARCHAR(255) NOT NULL,            -- 笔记标题
                       content TEXT NOT NULL,                  -- 笔记内容 (Markdown 格式)
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- 创建时间
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 更新时间
);
