import { Request, Response, Router } from 'express';
import Controller from '../interfaces/controller.interface';

class HelloWorldController implements Controller {
  public path = '/hello-world';
  public router = Router();

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.get(`${this.path}`, this.generateReport);
  }

  private generateReport = (request: Request, response: Response) => {
    response.send({
      mes: 'hello world',
    });
  }

}

export default HelloWorldController;
