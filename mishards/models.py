import logging
from sqlalchemy import (Integer, Boolean, Text,
        String, BigInteger, func, and_, or_,
        Column)
from sqlalchemy.orm import relationship, backref

from mishards import db

logger = logging.getLogger(__name__)

class TableFiles(db.Model):
    FILE_TYPE_NEW = 0
    FILE_TYPE_RAW = 1
    FILE_TYPE_TO_INDEX = 2
    FILE_TYPE_INDEX = 3
    FILE_TYPE_TO_DELETE = 4
    FILE_TYPE_NEW_MERGE = 5
    FILE_TYPE_NEW_INDEX = 6
    FILE_TYPE_BACKUP = 7

    __tablename__ = 'TableFiles'

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    table_id = Column(String(50))
    engine_type = Column(Integer)
    file_id = Column(String(50))
    file_type = Column(Integer)
    file_size = Column(Integer, default=0)
    row_count = Column(Integer, default=0)
    updated_time = Column(BigInteger)
    created_on = Column(BigInteger)
    date = Column(Integer)

    table = relationship(
        'Table',
        primaryjoin='and_(foreign(TableFile.table_id) == Table.table_id)',
        backref=backref('files', uselist=True, lazy='dynamic')
    )


class Tables(db.Model):
    TO_DELETE = 1
    NORMAL = 0

    __tablename__ = 'Tables'

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    table_id = Column(String(50), unique=True)
    state = Column(Integer)
    dimension = Column(Integer)
    created_on = Column(Integer)
    flag = Column(Integer, default=0)
    index_file_size = Column(Integer)
    engine_type = Column(Integer)
    nlist = Column(Integer)
    metric_type = Column(Integer)

    def files_to_search(self, date_range=None):
        cond = or_(
                TableFile.file_type==TableFile.FILE_TYPE_RAW,
                TableFile.file_type==TableFile.FILE_TYPE_TO_INDEX,
                TableFile.file_type==TableFile.FILE_TYPE_INDEX,
        )
        if date_range:
            cond = and_(
                cond,
                or_(
                    and_(TableFile.date>=d[0], TableFile.date<d[1]) for d in date_range
                    )
            )

        files = self.files.filter(cond)

        logger.debug('DATE_RANGE: {}'.format(date_range))
        return files
