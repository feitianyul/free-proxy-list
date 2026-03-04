**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-04 20:52:02 UTC（2026-03-05 04:52:02 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 14.56.177.44:3128 | ✓ 1714ms | ✓ 1372ms | ✓ 578ms | ✓ 1070ms | ✓ 891ms | http |
| 205.209.118.30:3138 | ✓ 570ms | ✓ 1186ms | ✓ 1236ms | ✓ 1375ms | 否 | http |
| 35.225.22.61:80 | ✓ 1678ms | ✓ 1924ms | ✓ 1530ms | ✓ 1144ms | ✓ 975ms | http |
| 125.128.12.144:3128 | ✓ 1650ms | ✓ 1137ms | ✓ 867ms | ✓ 981ms | ✓ 877ms | http |
| 61.72.221.194:3128 | ✓ 744ms | ✓ 1380ms | 否 | 否 | ✓ 1270ms | http |
| 14.56.107.244:3128 | ✓ 739ms | ✓ 1400ms | ✓ 1188ms | 否 | ✓ 1769ms | http |
| 125.128.12.14:3128 | ✓ 738ms | ✓ 1644ms | ✓ 823ms | ✓ 1002ms | ✓ 726ms | http |
| 121.128.121.54:3128 | ✓ 611ms | ✓ 1320ms | 否 | ✓ 918ms | ✓ 719ms | http |
| 61.72.221.94:3128 | ✓ 1190ms | ✓ 1818ms | 否 | 否 | ✓ 1094ms | http |
| 211.171.114.154:3128 | ✓ 853ms | ✓ 989ms | 否 | ✓ 1193ms | 否 | http |
| 61.72.110.94:3128 | ✓ 880ms | ✓ 1339ms | ✓ 1059ms | 否 | ✓ 1828ms | http |
| 217.76.245.80:999 | ✓ 897ms | 否 | ✓ 1176ms | ✓ 1678ms | ✓ 1412ms | http |
| 91.107.148.58:53967 | ✓ 689ms | ✓ 1917ms | 否 | 否 | ✓ 1857ms | http |
| 120.92.212.16:8890 | ✓ 1786ms | ✓ 1793ms | ✓ 1601ms | 否 | ✓ 916ms | http |
| 61.72.110.54:3128 | ✓ 1231ms | ✓ 1253ms | ✓ 821ms | ✓ 1364ms | ✓ 1577ms | http |
| 61.72.221.234:3128 | ✓ 1045ms | 否 | ✓ 749ms | ✓ 987ms | ✓ 1963ms | http |
| 81.70.169.194:80 | ✓ 920ms | ✓ 1283ms | ✓ 1039ms | ✓ 1154ms | ✓ 1045ms | http |
| 101.43.255.96:80 | ✓ 917ms | ✓ 1382ms | ✓ 1028ms | ✓ 1298ms | ✓ 944ms | http |
| 120.92.212.16:7890 | ✓ 936ms | ✓ 1127ms | ✓ 1600ms | 否 | ✓ 1186ms | http |
| 62.113.119.14:8080 | ✓ 876ms | 否 | ✓ 1111ms | 否 | ✓ 1267ms | http |
| 210.77.23.47:7897 | ✓ 796ms | ✓ 1754ms | ✓ 1072ms | ✓ 1609ms | ✓ 1118ms | http |
| 222.228.171.92:8080 | ✓ 796ms | 否 | ✓ 1634ms | ✓ 833ms | 否 | http |
| 103.215.36.88:16792 | ✓ 1113ms | ✓ 1366ms | ✓ 1090ms | ✓ 1345ms | ✓ 1012ms | http |
| 45.174.243.20:999 | ✓ 1143ms | ✓ 1625ms | ✓ 1422ms | 否 | 否 | http |
| 35.234.17.221:8080 | ✓ 1540ms | ✓ 1618ms | 否 | ✓ 1412ms | ✓ 846ms | http |
| 91.233.223.147:3128 | ✓ 1498ms | 否 | ✓ 1455ms | 否 | ✓ 1629ms | http |
| 103.215.36.88:16988 | ✓ 995ms | ✓ 1401ms | ✓ 1107ms | ✓ 1252ms | ✓ 1040ms | http |
| 165.227.5.10:8888 | ✓ 71ms | 否 | 否 | ✓ 1769ms | ✓ 511ms | http |
| 152.70.137.18:8888 | 否 | ✓ 1767ms | ✓ 1843ms | 否 | ✓ 1239ms | http |
| 103.139.138.194:3128 | ✓ 1864ms | 否 | ✓ 1570ms | ✓ 1498ms | ✓ 1088ms | http |
| 91.193.240.157:9877 | ✓ 1371ms | 否 | ✓ 1944ms | 否 | ✓ 1921ms | http |
| 136.226.48.24:11285 | ✓ 679ms | 否 | ✓ 819ms | ✓ 1581ms | ✓ 1300ms | http |
| 34.96.238.40:8080 | ✓ 1234ms | ✓ 1202ms | ✓ 1252ms | 否 | 否 | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1283ms | ✓ 1418ms | ✓ 1255ms | http |
| 121.230.8.251:1080 | ✓ 1087ms | ✓ 1171ms | ✓ 1013ms | ✓ 1400ms | ✓ 1031ms | http |
| 199.38.85.122:40004 | ✓ 1253ms | ✓ 1887ms | ✓ 1149ms | ✓ 1630ms | ✓ 1803ms | http |
| 199.38.85.122:40001 | ✓ 1254ms | 否 | ✓ 1632ms | ✓ 1690ms | ✓ 1281ms | http |
| 106.14.205.114:483 | ✓ 1760ms | ✓ 978ms | ✓ 1197ms | ✓ 1144ms | ✓ 1019ms | http |
| 5.75.196.26:40000 | ✓ 664ms | ✓ 1654ms | ✓ 1717ms | ✓ 1660ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1408ms | ✓ 1675ms | ✓ 848ms | ✓ 1354ms | ✓ 991ms | http |
| 121.43.196.210:8222 | ✓ 908ms | ✓ 1029ms | ✓ 838ms | ✓ 1047ms | ✓ 828ms | http |
| 121.43.196.213:8222 | ✓ 912ms | ✓ 1046ms | ✓ 874ms | ✓ 1090ms | ✓ 856ms | http |
| 114.55.226.123:10086 | ✓ 1057ms | ✓ 1315ms | ✓ 976ms | ✓ 1215ms | ✓ 959ms | http |
| 114.231.72.214:1080 | ✓ 938ms | ✓ 1141ms | ✓ 928ms | ✓ 1236ms | ✓ 948ms | http |
| 1.12.62.237:8080 | ✓ 1542ms | ✓ 1528ms | ✓ 1633ms | 否 | 否 | http |
| 20.120.225.109:3128 | 否 | ✓ 1036ms | ✓ 774ms | ✓ 830ms | ✓ 891ms | http |
| 121.230.8.61:1080 | ✓ 1478ms | ✓ 1691ms | ✓ 1577ms | 否 | ✓ 1653ms | http |
| 121.230.8.89:1080 | ✓ 1096ms | ✓ 1389ms | ✓ 944ms | ✓ 1447ms | 否 | http |
| 143.189.3.198:8080 | ✓ 520ms | ✓ 1922ms | ✓ 641ms | ✓ 809ms | ✓ 671ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1328ms | 否 | ✓ 1378ms | ✓ 1233ms | http |
| 46.249.103.192:443 | ✓ 809ms | 否 | ✓ 1537ms | 否 | ✓ 1712ms | http |
| 129.226.155.60:3128 | ✓ 1517ms | ✓ 1909ms | 否 | ✓ 1422ms | ✓ 1114ms | http |
| 210.223.44.230:3128 | ✓ 1811ms | ✓ 1041ms | ✓ 1848ms | ✓ 1151ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1856ms | 否 | 否 | ✓ 1423ms | ✓ 1475ms | http |
| 103.215.36.88:16206 | ✓ 1027ms | ✓ 1368ms | ✓ 982ms | ✓ 1215ms | ✓ 928ms | http |
| 107.155.65.87:13428 | 否 | 否 | ✓ 1044ms | ✓ 1022ms | ✓ 978ms | http |
| 101.255.208.17:8090 | ✓ 1666ms | 否 | 否 | ✓ 1623ms | ✓ 1712ms | http |
| 120.55.163.237:10086 | ✓ 846ms | ✓ 1012ms | ✓ 1524ms | ✓ 1058ms | ✓ 872ms | http |
| 45.140.147.82:1081 | 否 | 否 | ✓ 855ms | ✓ 1398ms | ✓ 1067ms | http |
| 90.84.188.97:8000 | ✓ 1032ms | 否 | 否 | ✓ 1998ms | ✓ 1787ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1315ms | 否 | ✓ 1324ms | ✓ 988ms | http |
| 120.92.211.211:7890 | ✓ 916ms | 否 | 否 | ✓ 1732ms | ✓ 1848ms | http |
| 45.140.147.82:1082 | ✓ 1577ms | ✓ 1694ms | ✓ 915ms | 否 | 否 | http |
| 200.125.171.254:999 | ✓ 1258ms | ✓ 1745ms | ✓ 1384ms | ✓ 1700ms | ✓ 1462ms | http |
| 138.124.53.25:7443 | ✓ 1102ms | 否 | ✓ 1806ms | 否 | ✓ 1697ms | http |
| 103.215.36.88:16894 | ✓ 927ms | ✓ 1230ms | ✓ 966ms | ✓ 1325ms | ✓ 1036ms | http |
| 103.215.36.88:10101 | ✓ 969ms | ✓ 1570ms | ✓ 1063ms | ✓ 1385ms | ✓ 967ms | http |
| 88.80.150.82:8080 | ✓ 1072ms | ✓ 1880ms | ✓ 1344ms | ✓ 1880ms | 否 | https |
| 103.215.36.88:17770 | ✓ 865ms | ✓ 1225ms | ✓ 962ms | 否 | ✓ 963ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
