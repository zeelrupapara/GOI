-- +migrate Down

-- Table: organization_collaborators
ALTER TABLE "organization_collaborators" DROP CONSTRAINT "fk_organization_collaborators_collaborator_id";
ALTER TABLE "organization_collaborators" DROP CONSTRAINT "fk_organization_collaborators_organization_id";

-- Table: assignees
ALTER TABLE "assignees" DROP CONSTRAINT "fk_assignees_collaborator_id";
ALTER TABLE "assignees" DROP CONSTRAINT "fk_assignees_activity_issue_id";
ALTER TABLE "assignees" DROP CONSTRAINT "fk_assignees_activity_pr_id";

-- Table: assigned_labals
ALTER TABLE "assigned_labals" DROP CONSTRAINT "fk_assigned_labals_labal_id";
ALTER TABLE "assigned_labals" DROP CONSTRAINT "fk_assigned_labals_activity_issue_id";
ALTER TABLE "assigned_labals" DROP CONSTRAINT "fk_assigned_labals_activity_pull_requests_id";

-- Table: reviews
ALTER TABLE "reviews" DROP CONSTRAINT "fk_reviews_reviewer_id";
ALTER TABLE "reviews" DROP CONSTRAINT "fk_reviews_pr_id";

-- Table: repositories
ALTER TABLE "repositories" DROP CONSTRAINT "fk_repositories_organization_collaborator_id";

-- Table: pull_requests
ALTER TABLE "pull_requests" DROP CONSTRAINT "fk_pull_requests_repository_id";
ALTER TABLE "pull_requests" DROP CONSTRAINT "fk_pull_requests_author_id";

-- Table: issues
ALTER TABLE "issues" DROP CONSTRAINT "fk_issues_repository_id";
ALTER TABLE "issues" DROP CONSTRAINT "fk_issues_author_id";

-- Table: commits
ALTER TABLE "commits" DROP CONSTRAINT "fk_commits_branch_id";
ALTER TABLE "commits" DROP CONSTRAINT "fk_commits_author_id";

-- Table: branches
ALTER TABLE "branches" DROP CONSTRAINT "fk_branches_repository_id";
