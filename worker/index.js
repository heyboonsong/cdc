const { Kafka } = require("kafkajs");

const kafka = new Kafka({
  clientId: "q0k00yjQRaqWmAAAZv955w",
  brokers: ["localhost:9092"],
});

const run = async () => {
  const consumer = kafka.consumer({
    groupId: "a898928a-7399-400d-9adc-0fc1ec8c5d83",
  });

  await consumer.connect();
  await consumer.subscribe({
    topic: "debezium-topic.public.users",
  });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      console.log({
        partition,
        offset: message.offset,
        value: message.value.toString(),
      });
    },
  });
};

run().catch(console.error);
