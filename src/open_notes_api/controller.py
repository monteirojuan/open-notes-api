"""Este módulo contém todas as ações que podem ser realizadas."""
from sqlalchemy import select
from sqlalchemy.ext.asyncio import async_sessionmaker

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
    """Lista todas as anotações."""
    async with db() as session:
        async with session.begin():
            notes = await session.execute(select(models.Note))
            return notes.scalars().all()


async def read_note(db: async_sessionmaker, note_id: int):
    """Retorna uma anotação."""
    async with db() as session:
        async with session.begin():
            notes = await session.execute(
                select(models.Note).where(models.Note.id == note_id)
            )
            return notes.scalars().one()


async def update_note(db: async_sessionmaker, note_id: int, data: schema.NoteUpdate):
    """Atualiza uma anotação."""
    async with db() as session:
        async with session.begin():
            query = await session.execute(
                select(models.Note).where(models.Note.id == note_id)
            )
            if note := query.scalars().first():
                if data.title:
                    note.title = data.title
                if data.content:
                    note.content = data.content
            await session.commit()
        await session.refresh(note)
        return note


async def delete_note(db: async_sessionmaker, note_id: int):
    """Deleta uma anotação."""
    async with db() as session:
        async with session.begin():
            query = await session.execute(
                select(models.Note).where(models.Note.id == note_id)
            )
            if note := query.scalars().first():
                await session.delete(note)
                await session.commit()
