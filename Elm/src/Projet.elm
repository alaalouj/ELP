module Projet exposing(..)
import Html exposing (..)
import Html.Attributes exposing(..)
import Html.Events exposing (onInput, onClick)
import String
import Browser
import Random
import Json.Decode exposing (map2, Decoder, field, int, string, map)
import Http 


-- Main 

main = 
  Browser.element
    { init = init
    , update = update 
    , subscriptions = subscriptions
    , view = view 
    }

-- Model

type alias Model =
  { etat : Etat
  , listeMots : List String
  , listeDef : (List Mots)
  , proposition : String}

type Etat
  = Failure
  | Loading
  | Success 

type alias Mots = 
  { word : String
  , meanings : (List Definition)}


type alias Definition =
  { partOfSpeech : String
  , definitions : (List Def)}

type alias Def = 
  { def : String}

init : () -> (Model, Cmd Msg)
init _ =
  ({etat = Loading ,listeMots=[],listeDef = [], proposition =""} , getListWord)

-- Update

type Msg 
  = MorePlease
  | GotWord (Result Http.Error String)
  | GotIndex Int
  | GotDef (Result Http.Error (List Mots))
  | Change String
  | GiveAnswer

update : Msg -> Model -> (Model, Cmd Msg)
update msg model = 
  case msg of 
    MorePlease ->
      ({ model |etat =Loading ,listeMots =[], listeDef = [], proposition = "" }, getListWord)
    GotWord result ->
      case result of
        Ok mots ->
          ( { model |etat =Success, listeMots = String.words mots, listeDef = [], proposition = "" }, Random.generate GotIndex (Random.int 1 1000) )
        Err _ ->
          ({ model |etat =Failure ,listeMots = [],listeDef = [], proposition = "" }, Cmd.none)
    GotIndex index ->
      case (getWordFromIndex model.listeMots index ) of 
        "Erreur" -> ({ model |etat =Failure ,listeMots = [],listeDef = [], proposition = "" }, Cmd.none)
        mot -> ( { model |etat =Success, listeMots = model.listeMots, listeDef = [], proposition = "" }, getRandomDef mot )
    GotDef result ->
      case result of 
        Ok def ->
          ({ model |etat =Success ,listeMots = model.listeMots,listeDef = def, proposition = "" }, Cmd.none)
        Err _ ->
          ({ model |etat =Failure ,listeMots = model.listeMots,listeDef = [], proposition = "" }, Cmd.none)
    Change try ->
      ({model | etat=model.etat,listeMots = model.listeMots,listeDef=model.listeDef, proposition =try}, Cmd.none)
    GiveAnswer ->
      ({model | etat = model.etat,listeMots = model.listeMots,listeDef=model.listeDef,proposition = leMot (model.listeDef)},Cmd.none)


-- SUBSCRIPTIONS 

subscriptions : Model -> Sub Msg
subscriptions model = 
  Sub.none


-- View 

view model = 
  div[] 
  [ h1 [] [text (if String.toLower (model.proposition) ==String.toLower(leMot( model.listeDef)) then leMot(model.listeDef) else "Guess the word !")]
  , viewWord model
  ]

viewWord : Model -> Html Msg
viewWord model = 
  case model.etat of 
    Failure ->
      div[]
        [ text "I could not load a word"
        , button [onClick MorePlease][text "Reload"]]
    Loading ->
      text "Loading ..."
    Success ->
      ul[]
        [ motsList model.listeDef  
        ,h3[][text "Type here : ", input [value model.proposition, onInput Change][],div[][text (if String.toLower(model.proposition) == String.toLower(leMot (model.listeDef)) then "Bien joué" else (if model.proposition == "" then "" else "Mauvaise réponse"))]
         ], button [ onClick MorePlease, style "display" "block" ][text "Play Again !"], button [onClick GiveAnswer, style "display" "block"][text "Show Answer"]]

  
getRandomDef : String -> Cmd Msg
getRandomDef mot = 
  Http.get 
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ mot
    , expect = Http.expectJson GotDef listMotsDecoder}

getListWord : Cmd Msg 
getListWord = 
  Http.get
    { url = "http://localhost:8000/listeMots.txt"
    , expect = Http.expectString GotWord}

getWordFromIndex : List String -> Int -> String
getWordFromIndex list index = 
  case list of 
    [] -> "Erreur"
    h::t -> 
      case index of 
        0 -> h
        _ -> getWordFromIndex t (index-1)

-- Fonctions 
defList : (List Def) -> Int -> Html Msg
defList definitions i = 
  case definitions of 
    [] -> div[][]
    (h::t) -> div [] [ul [] [ul[][text (String.fromInt(i)++ ". " ++ h.def)]], defList t (i+1)]
    
definitionList : (List Definition)-> Html Msg 
definitionList meanings = 
  case meanings of  
    [] -> div[][]
    (h::t) -> div[] [h3  [][ul [][text (h.partOfSpeech ++ " :")]], defList h.definitions 0,definitionList t ]

motsList : (List Mots) -> Html Msg
motsList mots =
  case mots of
    [] -> div[][]
    (h::t) -> div[][h2[][text "Meaning :" ],  definitionList h.meanings , motsList t]

leMot : (List Mots) -> String
leMot mots = 
  case mots of 
    [] -> "Probleme"
    h::_ -> h.word 



--Decoder 

motsDecoder : Decoder Mots
motsDecoder =
  map2 Mots
    (field "word" string)
    (field "meanings" listDefinitionDecoder)

listDefinitionDecoder :  Decoder (List Definition )
listDefinitionDecoder = Json.Decode.list definitionDecoder 

definitionDecoder : Decoder Definition
definitionDecoder =
  map2 Definition
    (field "partOfSpeech" string)
    (field "definitions" listDefDecoder)

listDefDecoder : Decoder (List Def)
listDefDecoder = Json.Decode.list defDecoder 

defDecoder : Decoder Def
defDecoder = 
  Json.Decode.map Def
    (field "definition" string)      

listMotsDecoder : Decoder (List Mots)
listMotsDecoder = Json.Decode.list motsDecoder


