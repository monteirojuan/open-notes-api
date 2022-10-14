from sqlalchemy.ext.asyncio import async_sessionmaker
from sqlalchemy import select

from open_notes_api import models


async def list_notes(db: async_sessionmaker):
    async with db() as session:
        async with session.begin():
            notes = await session.execute(select(models.Note))
    return notes.all()
