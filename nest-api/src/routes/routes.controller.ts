import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  Inject,
  OnModuleInit,
} from '@nestjs/common';
import { RoutesService } from './routes.service';
import { CreateRouteDto } from './dto/create-route.dto';
import { UpdateRouteDto } from './dto/update-route.dto';
import { ClientKafka, MessagePattern, Payload } from '@nestjs/microservices';
import { Producer } from 'kafkajs';

@Controller('routes')
export class RoutesController implements OnModuleInit {
  private kafkaProducer: Producer;

  constructor(
    private readonly routesService: RoutesService,
    @Inject('KAFKA_SERVICE')
    private readonly kafkaClient: ClientKafka,
  ) {}

  async onModuleInit() {
    this.kafkaProducer = await this.kafkaClient.connect();
  }

  @Post()
  create(@Body() createRouteDto: CreateRouteDto) {
    return this.routesService.create(createRouteDto);
  }

  @Get()
  findAll() {
    return this.routesService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.routesService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateRouteDto: UpdateRouteDto) {
    return this.routesService.update(+id, updateRouteDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.routesService.remove(+id);
  }

  @Get(':id/start')
  startRoute(@Param('id') id: string) {
    console.log(`Starting route ${id}`);
    this.kafkaProducer.send({
      topic: 'route.new-direction',
      messages: [
        {
          key: 'route.new-direction',
          value: JSON.stringify({
            id,
            clientId: '1',
          }),
        },
      ],
    });
  }

  @MessagePattern('route.new-position')
  consumeNewPosition(
    @Payload()
    message: {
      id: string;
      clientId: string;
      positions: [number, number];
      finished: boolean;
    },
  ) {
    console.log('New position received', message);
  }
}
