const types = require('../js/xdr_generated.js').default
const xdr = require('xdr-js-serialize')
const fs = require('fs')

const tx = new types.Transaction()
const action = new types.Action() 
const update = new types.Update()
update.contract(Buffer.from('3a547668e859fb7b112a1e2dd7efcb739176ab8cfd1d9f224847fce362ebd99c', 'hex'))
action.channelId(Buffer.from('3a547668e859fb7b112a1e2dd7efcb739176ab8cfd1d9f224847fce362ebd99c', 'hex'))
action.nonce(xdr.UnsignedHyper.fromXDR([0, 0, 0, 0, 0, 0, 0, 0]))
action.category(types.ActionCategory.update(update))
tx.signature(Buffer.from('33a547668e859fb7b112a1e2dd7efcb739176ab8cfd1d9f224847fce362ebd99ca547668e859fb7b112a1e2dd7efcb739176ab8cfd1d9f224847fce362ebd99c', 'hex'))
tx.address(Buffer.from('3a547668e859fb7b112a1e2dd7efcb739176ab8cfd1d9f224847fce362ebd99c', 'hex'))
tx.action(action)


fs.open('./examples/xdroutjs', 'w', function(err, fd) {
    if (err) {
        throw 'could not open file: ' + err;
    }

  const toWrite = tx.toXDR()
  for (let i = 0; i < 10000; i++) {
    // write the contents of the buffer, from position 0 to the end, to the file descriptor returned in opening our file
    fs.write(fd, toWrite, 0, toWrite.length, null, function(err) {
        if (err) throw 'error writing file: ' + err;
    });
  }
});
