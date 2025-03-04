-- 関連付けを削除
DELETE FROM label_tasks WHERE 
  (task_id, label_id) IN ((1,1), (1,2), (2,2), (2,3), (3,1), (3,3));

-- シードしたタスクを削除
DELETE FROM tasks WHERE id IN (1, 2, 3);

-- シードしたラベルを削除
DELETE FROM labels WHERE id IN (1, 2, 3);

-- シーケンスをリセット（オプション）
ALTER SEQUENCE tasks_id_seq RESTART WITH 1;
ALTER SEQUENCE labels_id_seq RESTART WITH 1;
