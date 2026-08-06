import {describe, it, beforeAll, afterAll, TestContext} from 'vitest';
import { ShellUse } from '@microsoft/shell-use';
import { afterEach } from 'vitest';

describe('Todo List Panel', () => {
  let su: ShellUse;

  // const printScreen = async (context: TestContext & Object): Promise<void>  =>{
  //   try {
  //     const screenText = await su.text();
  //     console.log(`\n=== SCREEN ON FAILURE (${context.task.name}) ===`);
  //     console.log(screenText);
  //     console.log('=== END ===\n');
  //   } catch {
  //     // session may already be closed
  //   }
  // }
  beforeAll(async () => {
    su = new ShellUse();
    await su.run('../../bin/tui-e2e-demo');
    await su.waitIdle();
    // Navigate to Todo List tab
    await su.press('3');
    await su.waitIdle();
  });

  afterAll(async () => {
    try {
      await su.close();
    } catch {
      // Session may already be closed
    }
  });

  // afterEach(async (context) => {
  //   if (context.task.result?.state === 'fail') {
  //     try {
  //       await printScreen(context);
  //     } catch {
  //       // session may already be closed
  //     }
  //   }
  // });

  it('adds a new todo item', async (testContext: TestContext & Object) => {
    // await printScreen(testContext)
    // Press 'a' to enter input mode
    await su.press('a');
    // await printScreen(testContext)
    await su.waitIdle();
    // Type the new item
    await su.type('New Todo Item');
    await su.waitIdle();
    // Press Enter to submit
    await su.press('Enter');
    await su.waitIdle();
    // Verify the new item appears
    await su.expectText('New Todo Item', {'timeout': 10 });
  });
});
