import datetime

from pydantic import BaseModel


class Note(BaseModel):
    """Schema utilizado na leitura das anotações."""
    id: int
    title: str
    content: str | None = None
    pinned: bool
    archived: bool
    created_at: datetime.datetime
    updated_at: datetime.datetime | None
    pinned_at: datetime.datetime | None
    archived_at: datetime.datetime | None

    class Config:
        orm_mode = True


class NoteCreate(BaseModel):
    """Schema utilizado na criação da anotação."""
    title: str
    content: str | None = None


class NoteUpdate(BaseModel):
    """Schema utilizado na atualização de uma anotação."""
    title: str | None = None
    content: str | None = None
