CREATE TABLE IF NOT EXISTS poles
(
    id         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    latitude   DECIMAL                   NOT NULL,
    longitude  DECIMAL                   NOT NULL,
    street     VARCHAR(255)              NOT NULL,
    bearing    ENUM ('w', 's', 'e', 'n') NOT NULL,
    kind       VARCHAR(255),
    height     VARCHAR(6)                NOT NULL,
    locked     TINYINT                   NOT NULL DEFAULT 0,
    created_by VARCHAR(255)              NOT NULL,
    updated_by VARCHAR(255)              NOT NULL,
    created_at TIMESTAMP                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
)
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS pole_installations
(
    id         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    pole_id    INT UNSIGNED  NOT NULL,
    ubihub_sn  VARCHAR(25)   NOT NULL,
    cameras_sn VARCHAR(2048) NOT NULL,
    start      TIMESTAMP,
    end        TIMESTAMP,
    created_by VARCHAR(255)  NOT NULL,
    updated_by VARCHAR(255)  NOT NULL,
    created_at TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT pole_installations_pole_id_fk FOREIGN KEY (pole_id) REFERENCES poles (id)
)
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS pole_audits
(
    id                   INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    pole_installation_id INT UNSIGNED                                            NOT NULL,
    attempt              INT UNSIGNED                                            NOT NULL,
    state                ENUM ('pending', 'approved', 'remediation', 'rejected') NOT NULL DEFAULT 'pending',
    summary              VARCHAR(2048),
    auditor              VARCHAR(255)                                            NOT NULL,
    created_by           VARCHAR(255)                                            NOT NULL,
    updated_by           VARCHAR(255)                                            NOT NULL,
    created_at           TIMESTAMP                                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP                                               NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT pole_audits_pole_installation_id_fk FOREIGN KEY (pole_installation_id) REFERENCES pole_installations (id),
    CONSTRAINT pole_audits_unique_idx UNIQUE (pole_installation_id, attempt)
)
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS pole_audit_notes
(
    id            INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    pole_audit_id INT UNSIGNED                             NOT NULL,
    type          ENUM ('photo', 'audio', 'video', 'text') NOT NULL DEFAULT 'photo',
    datum         VARCHAR(4096)                            NOT NULL,
    created_by    VARCHAR(255)                             NOT NULL,
    updated_by    VARCHAR(255)                             NOT NULL,
    created_at    TIMESTAMP                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP                                NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT pole_audit_notes_pole_audits_id_fk FOREIGN KEY (pole_audit_id) REFERENCES pole_audits (id)
)
    COLLATE = utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS pole_audit_questions
(
    id         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    device     ENUM ('cell', 'dtm', 'tvm', 'hub')                      NOT NULL,
    position   INT UNSIGNED                                            NOT NULL,
    question   VARCHAR(1024)                                           NOT NULL,
    input      ENUM ('radio', 'text', 'selections', 'number', 'photo') NOT NULL,
    answer     VARCHAR(4096)                                           NOT NULL,
    created_by VARCHAR(255)                                            NOT NULL,
    updated_by VARCHAR(255)                                            NOT NULL,
    created_at TIMESTAMP                                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
)
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS pole_audit_question_answers
(
    id         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    datum      VARCHAR(4096) NOT NULL,
    created_by VARCHAR(255)  NOT NULL,
    created_at TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
)
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS pole_auditables
(
    pole_audit_id                 INT UNSIGNED NOT NULL,
    pole_audit_question_id        INT UNSIGNED NOT NULL,
    pole_audit_question_answer_id INT UNSIGNED,

    CONSTRAINT pole_auditables_pole_audit_id_fk FOREIGN KEY (pole_audit_id) REFERENCES pole_audits (id),
    CONSTRAINT pole_auditables_pole_audit_question_id_fk FOREIGN KEY (pole_audit_question_id) REFERENCES pole_audit_questions (id),
    CONSTRAINT pole_auditables_pole_audit_question_answer_id_fk FOREIGN KEY (pole_audit_question_answer_id) REFERENCES pole_audit_question_answers (id)
)
    COLLATE = utf8mb4_unicode_ci;