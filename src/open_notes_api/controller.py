from sqlalchemy.ext.asyncio import async_sessionmaker
from sqlalchemy import select

from open_notes_api import models, schema


async def create_note(db: async_sessionmaker, note: schema.NoteCreate):
    """Realiza a criação de uma anotação no banco de dados."""
    async with db() as session:
        async with session.begin():
            db_note = models.Note(
                title=note.title,
                content=note.content,
            )
            session.add(db_note)
            await session.commit()

        await session.refresh(db_note)
    return db_note


async def list_notes(db: async_sessionmaker):
    async with db() as session:
        async with session.begin():
            notes = await session.execute(select(models.Note))
            return notes.scalars().all()
