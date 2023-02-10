import { Prop, raw, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

@Schema()
export class Route {
  @Prop()
  title: string;

  @Prop(
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  startPosition: Position;

  @Prop(
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  endPosition: Position;
}

export type Position = {
  lat: number;
  lng: number;
};

export type RouteDocument = Route & Document;

export const RouteSchema = SchemaFactory.createForClass(Route);
