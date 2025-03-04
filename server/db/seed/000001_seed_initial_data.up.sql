-- タスクデータの挿入
INSERT INTO tasks (id, title, description, priority, status, started_at, finished_at, created_at, updated_at)
VALUES
  (1, 'Task1', 'Description1', 'Low', 'Pending', now(), now(), now(), now()),
  (2, 'Task2', 'Description2', 'Middle', 'In_Progress', now(), now(), now(), now()),
  (3, 'Task3', 'Description3', 'High', 'Completed', now(), now(), now(), now())
ON CONFLICT (id)
DO UPDATE SET
    title = excluded.title,
    description = excluded.description,
    priority = excluded.priority,
    status = excluded.status,
    started_at = excluded.started_at,
    finished_at = excluded.finished_at,
    created_at = excluded.created_at,
    updated_at = excluded.updated_at;
    
-- シーケンスをリセット
SELECT pg_catalog.setval('tasks_id_seq', (select max(id) from tasks), true);

-- ラベルデータの挿入
INSERT INTO labels (id, name, created_at, updated_at)
VALUES
  (1, 'Label1', now(), now()),
  (2, 'Label2', now(), now()),
  (3, 'Label3', now(), now())
ON CONFLICT (id)
DO UPDATE SET
    name = excluded.name,
    created_at = excluded.created_at,
    updated_at = excluded.updated_at;
    
-- シーケンスをリセット
SELECT pg_catalog.setval('labels_id_seq', (select max(id) from labels), true);

-- タスクとラベルの関連付け
INSERT INTO label_tasks (task_id, label_id)
VALUES
  (1, 1),
  (1, 2),
  (2, 2),
  (2, 3),
  (3, 1),
  (3, 3)
ON CONFLICT (task_id, label_id) DO NOTHING;
