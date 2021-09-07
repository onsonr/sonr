process.env.BABEL_ENV = 'development';
process.env.NODE_ENV = 'development';

process.on('unhandledRejection', (err) => {
  console.error(err);
  throw err;
});

const chalk = require('chalk');
const fs = require('fs-extra');
const webpack = require('webpack');
const clearConsole = require('react-dev-utils/clearConsole');
const formatWebpackMessages = require('react-dev-utils/formatWebpackMessages');
const forkTsCheckerWebpackPlugin = require('react-dev-utils/ForkTsCheckerWebpackPlugin');
const typescriptFormatter = require('react-dev-utils/typescriptFormatter');
const checkRequiredFiles = require('react-dev-utils/checkRequiredFiles');

const paths = require('../config/paths');
const webpackConfig = require('../config/webpack.config');
const appName = require(paths.appPackageJson).name;

const config = webpackConfig('development');
const isInteractive = process.stdout.isTTY;
const tscCompileOnError = process.env.TSC_COMPILE_ON_ERROR === 'true';
const useTypeScript = fs.existsSync(paths.appTsConfig);

function printInstructions(appName) {
  console.log();
  console.log(
    `You can now load ${chalk.bold(appName)} from ${chalk.bold(
      'chrome://extensions'
    )}.`
  );
  console.log(
    `Remember to select the ${chalk.bold(
      '/build'
    )} folder when loading the unpacked extension.`
  );
  console.log();

  console.log();
  console.log('Note that the development build is not optimized.');
  console.log(
    'Use "npm run assemble" to build and ZIP your extension for production.'
  );
  console.log();
}

// Warn and crash if required files are missing
if (!checkRequiredFiles([paths.appHtml, paths.appIndexJs])) {
  process.exit(1);
}

let compiler;
try {
  compiler = webpack(config);
} catch (err) {
  console.log('Failed to compile.');
  console.log();
  console.log(err.message || err);
  console.log();
  process.exit(1);
}

compiler.hooks.invalid.tap('invalid', () => {
  if (isInteractive) {
    clearConsole();
  }
  console.log('Compiling...');
});

let isFirstCompile = true;
let tsMessagesPromise;
let tsMessagesResolver;

if (useTypeScript) {
  compiler.hooks.beforeCompile.tap('beforeCompile', () => {
    tsMessagesPromise = new Promise((resolve) => {
      tsMessagesResolver = (msgs) => resolve(msgs);
    });
  });

  forkTsCheckerWebpackPlugin
    .getCompilerHooks(compiler)
    .receive.tap('afterTypeScriptCheck', (diagnostics, lints) => {
      const allMsgs = [...diagnostics, ...lints];
      const format = (message) =>
        `${message.file}\n${typescriptFormatter(message, true)}`;

      tsMessagesResolver({
        errors: allMsgs.filter((msg) => msg.severity === 'error').map(format),
        warnings: allMsgs
          .filter((msg) => msg.severity === 'warning')
          .map(format),
      });
    });
}

compiler.hooks.done.tap('done', async (stats) => {
  if (isInteractive) {
    clearConsole();
  }

  const statsData = stats.toJson({
    all: false,
    warnings: true,
    errors: true,
  });

  if (useTypeScript && statsData.errors.length === 0) {
    const delayedMsg = setTimeout(() => {
      console.log(
        chalk.yellow(
          'Files successfully emitted, waiting for typecheck results...'
        )
      );
    }, 100);

    const messages = await tsMessagesPromise;
    clearTimeout(delayedMsg);
    if (tscCompileOnError) {
      statsData.warnings.push(...messages.errors);
    } else {
      statsData.errors.push(...messages.errors);
    }
    statsData.warnings.push(...messages.warnings);

    if (tscCompileOnError) {
      stats.compilation.warnings.push(...messages.errors);
    } else {
      stats.compilation.errors.push(...messages.errors);
    }
    stats.compilation.warnings.push(...messages.warnings);

    if (messages.errors.length > 0) {
      if (tscCompileOnError) {
        devSocket.warnings(messages.errors);
      } else {
        devSocket.errors(messages.errors);
      }
    } else if (messages.warnings.length > 0) {
      devSocket.warnings(messages.warnings);
    }

    if (isInteractive) {
      clearConsole();
    }
  }

  const messages = formatWebpackMessages(statsData);
  const isSuccessful = !messages.errors.length && !messages.warnings.length;
  if (isSuccessful) {
    console.log(chalk.green('Compiled successfully!'));
  }
  if (isSuccessful && (isInteractive || isFirstCompile)) {
    printInstructions(appName);
  }
  isFirstCompile = false;

  if (messages.errors.length) {
    if (messages.errors.length > 1) {
      messages.errors.length = 1;
    }
    console.log(chalk.red('Failed to compile.\n'));
    console.log(messages.errors.join('\n\n'));
    return;
  }

  if (messages.warnings.length) {
    console.log(chalk.yellow('Compiled with warnings.\n'));
    console.log(messages.warnings.join('\n\n'));

    console.log(
      '\nSearch for the ' +
        chalk.underline(chalk.yellow('keywords')) +
        ' to learn more about each warning.'
    );
    console.log(
      'To ignore, add ' +
        chalk.cyan('// eslint-disable-next-line') +
        ' to the line before.\n'
    );
  }
});

compiler.watch({}, function (err) {
  if (err) {
    console.error(err);
  }
});
