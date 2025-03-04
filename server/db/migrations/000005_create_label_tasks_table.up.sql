CREATE TABLE "label_tasks" (
    "label_id" bigint NOT NULL, 
    "task_id" bigint NOT NULL, 
    "created_at" timestamptz NOT NULL  DEFAULT now(), 
    "updated_at" timestamptz NOT NULL  DEFAULT now(), 
    PRIMARY KEY ("label_id", "task_id"), 
    CONSTRAINT "label_tasks_label_id" FOREIGN KEY ("label_id") REFERENCES "labels" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, 
    CONSTRAINT "label_tasks_task_id" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
