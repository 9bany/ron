import 'dotenv/config';
import HelloWorldController from './hello/hello.controller';
import * as bodyParser from 'body-parser';
import * as cookieParser from 'cookie-parser';
import * as express from 'express';
import Controller from './interfaces/controller.interface';

class App {
  public app: express.Application;

  constructor(controllers: Controller[]) {
    this.app = express();
    this.initializeMiddlewares();
    this.initializeControllers(controllers);
  }

  public listen() {
    this.app.listen(process.env.PORT, () => {
      console.log('Listening on port: ' + process.env.PORT);
    });
  }

  public getServer() {
    return this.app;
  }

  private initializeMiddlewares() {
    this.app.use(bodyParser.json());
    this.app.use(cookieParser());
  }

  private initializeControllers(controllers: Controller[]) {
    console.log(controllers);
    controllers.forEach((controller) => {
      this.app.use('/', controller.router);
    });
  }

}

const app = new App(
  [
    new HelloWorldController(),
  ],
);

app.listen();
