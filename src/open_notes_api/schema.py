import datetime

from pydantic import BaseModel


class Note(BaseModel):
    id: int
    title: str
    content: str | None = None
    pinned: bool
    archived: bool
    created_at: datetime.datetime
    updated_at: datetime.datetime
    pinned_at: datetime.datetime
    archived_at: datetime.datetime
