-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS philosophy_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    who_is_philosopher TEXT NOT NULL, -- каков портрет философа / мыслителя?
    themes TEXT NOT NULL, -- темы, на которых хотелось бы сосредоточиться
    disciplines TEXT NOT NULL, -- философские дисциплины, на которых хотелось бы сосредоточиться
    three_philosophers TEXT NOT NULL, -- три философа, которые сразу пришли на ум
    three_texts TEXT NOT NULL, -- три философских текста, которые пользователь читал либо просто смог вспомнить

    whoami TEXT NOT NULL, -- самоописание, портрет, ответ на вопрос "какое из следующих высказываний описывает вас лучше всего"
    motivation TEXT NOT NULL, -- мотивация изучать философию
    philosophy_assoc_era TEXT NOT NULL, -- с какой эпохой больше всего ассоциируется философия
    favorite_era TEXT NOT NULL, -- эпоха, в которой есть найбольшая заинтересованность / с которой человек себя ассоциирует

    preferred_text_size TEXT NOT NULL, -- предпочтительный размер текстов: small, medium, large
    is_interested_in_east BOOLEAN DEFAULT FALSE, -- заинтересованность в восточной философии
    is_interested_in_media BOOLEAN DEFAULT FALSE, -- заинтересованность в видео, подкастах, фильмах, а не только в книгах
    is_interested_in_fiction BOOLEAN DEFAULT FALSE, -- заинтересованность в художественной литературе, а не только в строго философских
    is_interested_in_metaphorical BOOLEAN DEFAULT FALSE, -- заинтересованность не только в строгой, последовательной и систематической, но и в мистической, метафорической, литературной философии
    is_interested_in_rare_voices BOOLEAN DEFAULT FALSE, -- заинтересованность в изучении голосов, которые редко упоминаются в "официальной" истории философии: философини, забытые мыслители, "фон" эпохи и т.д.
    is_interested_in_state_of_the_art BOOLEAN DEFAULT FALSE, -- заинтересованность в самых современных философских проблемах и дискуссиях

    other_interests TEXT, -- другие интересы, которые есть у человека и с помощью которых можно дополнить изучение философии

    open_comment TEXT, -- открытый комментарий: всё, что хотелось бы добавить

    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS philosophy_requests;
-- +goose StatementEnd
