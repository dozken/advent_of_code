export class HouseBuilder {
  constructor(builderfunction: (builder: HouseBuilder) => void = (e) => e) {
    builderfunction(this);
  }

  build(): House {
    return {} as House
  }

  private name: string = "Grape House";

  withName(name: string): this {
    this.name = name;
    return this;
  }

  private roofBuilder: RoofBuilder = new RoofBuilder();

  withRoof(
    builderFunction: (builder: RoofBuilder) => void = (e) => e
  ): this {
    builderFunction(this.roofBuilder);
    return this;
  }

  private roomBuilders: RoomBuilder[] = [];

  withRoom(builderFunction: (builder: RoomBuilder) => void = (e) => e): this {
    const newBuiler = new RoomBuilder();
    builderFunction(newBuiler);
    this.roomBuilders.push(newBuiler);
    return this;
  }
}
