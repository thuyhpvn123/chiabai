var output = document.getElementById("log-content");
var socket = new WebSocket("ws://127.0.0.1:2000/ws");
var socketActive = false;
var $createResult = document.getElementById('create-result');

var deck ;
console.log("Imported");
// * Websocket
// Connect to server successfully
var walletAddress = "aa39344b158f4004cac70bb4ace871a9b54baa1e";

var messageForm = {
  command:"get-wallet-pagination",
  limit:2,
  page:1,
};

socket.onopen = (msg) => {
  socketActive = true;


  //Send walletAddress to server
  let walletMessage = structuredClone(messageForm);
  walletMessage.type = "WalletMessage";
  walletMessage.message = walletAddress;
  // sendMessage(walletMessage);
};

// WS connection's closed
socket.onclose = (event) => {
  console.log("WS Connection is closed: ", event);
};

// WS connection having errors
socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};
socket.onmessage = (msg) => {
  var data12 = JSON.parse(msg.data);
  output.innerHTML += "Server: " + msg.data + "\n";
   switch (data12.command) {
    case "generate-keys":
        console.log("data12.value:",data12.data)
        document.getElementById('keys').innerHTML = data12.data.message
        break;
    case "encrypt-cards":
      console.log("data12.value:",data12.data)
      deck = data12.data.message
      // dataCall.inputArray = deck
      console.log("deck:",deck)
      break;
  }

}
var sendMessage = (msg) => {
  console.log(msg);
  socket.send(JSON.stringify(msg));

};
// connect wallet
var $connectNode = document.getElementById('register');

$connectNode.addEventListener('submit', async(e) => {

    e.preventDefault()
    console.log("connect node")
    var flag =1,
    fromAddress = $('#address').val()
    prikey = $('#prikey').val()

    if( fromAddress ==''){
      flag=0
      $('.error_address').html("Please type address of the account")
    }else{
      $('.error_address').html("")
    }
    if( prikey ==''){
      flag=0
      $('.error_prikey').html("Please type private key of the account")
    }else{
      $('.error_prikey').html("")
    }

    if(flag==1  ){
      try{
        await connectWallet(fromAddress,prikey);
        console.log("connected wallet ")
        $createResult.innerHTML =` connect wallet ${fromAddress} successfully` ;
        $('#wallet-id-user').html(`${fromAddress}`) ;
      }catch{
        console.log(e)
        $createResult.innerHTML = `Ooops... there was an error while trying to register wallet`;
      }
    }
  })
 

var connectWallet = (address,priKey) => {
    var wallet={
        "address":address,
        "priKey":priKey
      }
    var connectWallet = {
        command: "connect-wallet",
        value: wallet,  
    }
    sendMessage(connectWallet);
    console.log("connectWallet")
};

//generate keys
var $generateKeys = document.getElementById('generateKey');

$generateKeys.addEventListener('click', async(e) => {

    e.preventDefault()
    console.log("generate keys")
      try{
        await generate(4);
        console.log("generated keys ")
      }catch{
        console.log(e)
        $createResult.innerHTML = `Ooops... there was an error while trying to register wallet`;
      }
})
var generate = (number) => {
    var ms={
        "numPlayers":number
      }
    var generateMsg = {
        command: "generate-keys",
        value: ms,  
    }
    sendMessage(generateMsg);
    console.log("generate-keys")
};
//shuffle deck
var shuffleDeck = (key1,key2,key3,key4) => {
    var playerKeys =[key1,key2,key3,key4]
    var ms={
        "playerKeys":playerKeys
      }

    var shufferMsg = {
        command: "shuffle-and-encrypt-cards",
        value: ms,  
    }
    sendMessage(shufferMsg);
    console.log("shuffle cards")
};
var $encrypt = document.getElementById('encrypt');

$encrypt.addEventListener('submit', async(e) => {

    e.preventDefault()
    console.log("shuffle cards")
    var flag =1,key1,key2,key3,key4
    key1 = $('#key1').val()
    key2 = $('#key2').val()
    key3 = $('#key3').val()
    key4 = $('#key4').val()

    if( key1 ==''|| key2 ==''|| key3 ==''|| key4 ==''){
      flag=0
      $('.error_key').html("Please type key")
    }else{
      $('.error_key').html("")
    }
    if(flag==1  ){
      try{
        await shuffleDeck(key1,key2,key3,key4);
        console.log("shuffle and encrypted cards successfully ")
      }catch{
        console.log(e)
        $createResult.innerHTML = `Ooops... there was an error while trying to register wallet`;
      }
    }
  })
  //call function
  // var Inputs=JSON.stringify(deck)
var $dealCards = document.getElementById('dealCards');

$dealCards.addEventListener('click', async(e) => {

    e.preventDefault()
      try{
        console.log("deal cards")
        var inputs=JSON.stringify(
          {
            "internalType": "string[]",
            "name": "cardsArr",
            "type": "string[]",
            "value": deck,
        
          });
          var functionInputs=[inputs] ;

          var dataCall = 
          {
            'from':   "45c75cfb8e20a8631c134555fa5d61fcf3e602f2",
            'priKey': "36e1aa979f98c7154fb2491491ec044ccac099651209ccfbe2561746dbe29ebb",
            'to':    "f8eaba3eb679f6defbe78ce8dd5229ec3622f2a7",
            amount:            "",
            "function-name":"setDeck",
            inputArray:functionInputs,
            gas:1000000,
            gasPrice:10,
            timeUse:1000,
            relatedAddresses:[],
          }
        ;
        await deal(dataCall);
        console.log("deal cards")
      }catch{
        console.log(e)
        $createResult.innerHTML = `Ooops... there was an error while trying to register wallet`;
      }
})
var deal = (ms) => {

    var dealMsg = {
        command: "send-transaction",
        value: ms,  
    }
    console.log("deck nè: ",deck)
    console.log("ms nè: ",ms)
    sendMessage(dealMsg);
    console.log("deal-cards")
};

//test
var $test = document.getElementById('test');

$test.addEventListener('click', async(e) => {

    e.preventDefault()
      try{
        console.log("test")
        var in2=JSON.stringify(
          {
            "internalType": "uint256",
            "name": "amountOut",
            "type": "uint256",
            "value": 9,
        
          },
          {
            "internalType": "address[]",
            "name": "path",
            "type": "address[]",
            "value": 9,
        
          },
          {
            "internalType": "address",
            "name": "to",
            "type": "address",
            "value": 9,
        
          },
          {
            "internalType": "uint256",
            "name": "deadline",
            "type": "uint256",
            "value": 9,
        
          }
        );
           var functionInputs=[in2] ;

          var dataCall = 
          {
            'from':   "45c75cfb8e20a8631c134555fa5d61fcf3e602f2",
            'priKey': "36e1aa979f98c7154fb2491491ec044ccac099651209ccfbe2561746dbe29ebb",
            'to':    "1f171bf125864bcd00e8622c44030b39843f31f9",
            amount:            "",
            "function-name":"swapETHForExactTokens",
            inputArray:functionInputs,
            gas:1000000,
            gasPrice:10,
            timeUse:1000,
            relatedAddresses:[],
          }
        ;
        await deal(dataCall);
        console.log("deal cards")
      }catch{
        console.log(e)
        $createResult.innerHTML = `Ooops... there was an error while trying to register wallet`;
      }
})
var deal = (ms) => {

    var dealMsg = {
        command: "send-transaction",
        value: ms,  
    }
    console.log("deck nè: ",deck)
    console.log("ms nè: ",ms)
    sendMessage(dealMsg);
    console.log("deal-cards")
};