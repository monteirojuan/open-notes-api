from sqlalchemy.orm import DeclarativeBase, mapped_column
from sqlalchemy import Integer, DateTime, String, Text, Boolean, func, false


class Base(DeclarativeBase):
    pass


class Note(Base):
    __tablename__ = "notes"

    id = mapped_column(Integer, primary_key=True)
    title = mapped_column(String)
    content = mapped_column(Text)
    pinned = mapped_column(Boolean, default=false())
    archived = mapped_column(Boolean, default=false())
    created_at = mapped_column(DateTime, server_default=func.now())
    updated_at = mapped_column(DateTime, server_default=func.now(), server_onupdate=func.now())
    pinned_at = mapped_column(DateTime)
    archived_at = mapped_column(DateTime)

    __mapper_args__ = {"eager_defaults": True}
