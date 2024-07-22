FROM google/cloud-sdk:latest
RUN curl -O https://fastdl.mongodb.org/tools/db/mongodb-database-tools-debian11-x86_64-100.8.0.deb && \
    apt-get install ./mongodb-database-tools-debian11-x86_64-100.8.0.deb && \
    rm -rf ./mongodb-database-tools-debian11-x86_64-100.8.0.deb