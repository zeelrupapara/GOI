-- +migrate Up

-- Table: organization_collaborators
ALTER TABLE "organization_collaborators" ADD CONSTRAINT "fk_organization_collaborators_collaborator_id" FOREIGN KEY ("collaborator_id") REFERENCES "collaborators" ("id");
ALTER TABLE "organization_collaborators" ADD CONSTRAINT "fk_organization_collaborators_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");

-- Table: assignees
ALTER TABLE "assignees" ADD CONSTRAINT "fk_assignees_collaborator_id" FOREIGN KEY ("collaborator_id") REFERENCES "collaborators" ("id");
ALTER TABLE "assignees" ADD CONSTRAINT "fk_assignees_activity_issue_id" FOREIGN KEY ("activity_id") REFERENCES "issues" ("id");
ALTER TABLE "assignees" ADD CONSTRAINT "fk_assignees_activity_pr_id" FOREIGN KEY ("activity_id") REFERENCES "pull_requests" ("id");

-- Table: assigned_labals
ALTER TABLE "assigned_labals" ADD CONSTRAINT "fk_assigned_labals_labal_id" FOREIGN KEY ("labal_id") REFERENCES "labals" ("id");
ALTER TABLE "assigned_labals" ADD CONSTRAINT "fk_assigned_labals_activity_issue_id" FOREIGN KEY ("activity_id") REFERENCES "issues" ("id");
ALTER TABLE "assigned_labals" ADD CONSTRAINT "fk_assigned_labals_activity_pull_requests_id" FOREIGN KEY ("activity_id") REFERENCES "pull_requests" ("id");

-- Table: reviews
ALTER TABLE "reviews" ADD CONSTRAINT "fk_reviews_reviewer_id" FOREIGN KEY ("reviewer_id") REFERENCES "collaborators" ("id");
ALTER TABLE "reviews" ADD CONSTRAINT "fk_reviews_pr_id" FOREIGN KEY ("pr_id") REFERENCES "pull_requests" ("id");

-- Table: repositories
ALTER TABLE "repositories" ADD CONSTRAINT "fk_repositories_organization_collaborator_id" FOREIGN KEY ("organization_collaborator_id") REFERENCES "organization_collaborators" ("id");

-- Table: pull_requests
ALTER TABLE "pull_requests" ADD CONSTRAINT "fk_pull_requests_repository_id" FOREIGN KEY ("repository_id") REFERENCES "repositories" ("id");
ALTER TABLE "pull_requests" ADD CONSTRAINT "fk_pull_requests_author_id" FOREIGN KEY ("author_id") REFERENCES "collaborators" ("id");

-- Table: issues
ALTER TABLE "issues" ADD CONSTRAINT "fk_issues_repository_id" FOREIGN KEY ("repository_id") REFERENCES "repositories" ("id");
ALTER TABLE "issues" ADD CONSTRAINT "fk_issues_author_id" FOREIGN KEY ("author_id") REFERENCES "collaborators" ("id");

-- Table: commits
ALTER TABLE "commits" ADD CONSTRAINT "fk_commits_author_id" FOREIGN KEY ("author_id") REFERENCES "collaborators" ("id");
ALTER TABLE "commits" ADD CONSTRAINT "fk_commits_branch_id" FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

-- Table: branches
ALTER TABLE "branches" ADD CONSTRAINT "fk_branches_repository_id" FOREIGN KEY ("repository_id") REFERENCES "repositories" ("id");
