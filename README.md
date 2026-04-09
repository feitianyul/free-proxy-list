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

最后更新：2026-04-09 11:59:01 UTC（2026-04-09 19:59:01 UTC+8）

**代理总数：339**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 339 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 339 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.57:8448 | ✓ 824ms | ✓ 1312ms | ✓ 1189ms | ✓ 1460ms | ✓ 521ms | http |
| 8.209.239.31:30000 | ✓ 455ms | 否 | ✓ 521ms | ✓ 1086ms | ✓ 550ms | http |
| 147.161.210.140:8800 | ✓ 779ms | 否 | ✓ 712ms | ✓ 936ms | ✓ 790ms | http |
| 154.40.137.209:55965 | ✓ 1107ms | ✓ 1632ms | ✓ 1460ms | ✓ 1265ms | ✓ 1032ms | http |
| 1.231.81.166:3128 | ✓ 785ms | ✓ 993ms | ✓ 1717ms | ✓ 1183ms | ✓ 743ms | http |
| 35.225.22.61:80 | ✓ 861ms | 否 | ✓ 1125ms | ✓ 1736ms | ✓ 960ms | http |
| 113.160.132.26:8080 | ✓ 800ms | 否 | ✓ 1693ms | ✓ 1167ms | ✓ 851ms | http |
| 167.103.115.102:8800 | ✓ 1406ms | 否 | ✓ 967ms | ✓ 1226ms | ✓ 1124ms | http |
| 5.104.87.17:8051 | ✓ 1238ms | 否 | ✓ 1329ms | ✓ 1514ms | ✓ 1595ms | http |
| 217.174.244.117:3129 | ✓ 742ms | 否 | 否 | ✓ 1820ms | ✓ 1546ms | http |
| 167.103.34.108:8800 | ✓ 1476ms | 否 | ✓ 1396ms | ✓ 1625ms | ✓ 1381ms | http |
| 116.80.63.178:3172 | ✓ 1740ms | 否 | 否 | ✓ 1950ms | ✓ 1823ms | http |
| 38.145.208.244:8446 | ✓ 159ms | ✓ 1079ms | ✓ 460ms | ✓ 698ms | ✓ 508ms | http |
| 38.92.10.139:33985 | ✓ 290ms | ✓ 756ms | ✓ 726ms | ✓ 843ms | ✓ 542ms | http |
| 38.145.220.49:8444 | ✓ 214ms | ✓ 1758ms | ✓ 761ms | ✓ 655ms | ✓ 529ms | http |
| 38.145.218.234:8447 | ✓ 309ms | ✓ 1192ms | ✓ 83ms | ✓ 680ms | ✓ 646ms | http |
| 38.145.218.229:8444 | ✓ 411ms | ✓ 1187ms | ✓ 139ms | ✓ 677ms | ✓ 768ms | http |
| 152.32.132.190:7890 | ✓ 835ms | ✓ 1316ms | ✓ 905ms | ✓ 794ms | ✓ 658ms | http |
| 167.103.144.127:8800 | ✓ 1533ms | 否 | ✓ 927ms | ✓ 1663ms | ✓ 1676ms | http |
| 45.12.151.226:2829 | ✓ 1253ms | 否 | ✓ 1341ms | ✓ 1761ms | 否 | http |
| 161.35.70.36:8888 | ✓ 1061ms | 否 | ✓ 1981ms | 否 | ✓ 1496ms | http |
| 167.103.31.122:8800 | ✓ 1643ms | 否 | ✓ 1842ms | 否 | ✓ 1558ms | http |
| 120.92.212.16:7890 | ✓ 1885ms | 否 | ✓ 888ms | ✓ 1173ms | ✓ 1413ms | http |
| 59.46.216.131:30001 | ✓ 1998ms | ✓ 1265ms | 否 | ✓ 1251ms | ✓ 1083ms | http |
| 45.167.124.52:8080 | ✓ 905ms | 否 | ✓ 701ms | ✓ 1767ms | 否 | http |
| 45.167.125.21:999 | ✓ 1180ms | 否 | ✓ 1373ms | ✓ 1716ms | ✓ 1521ms | http |
| 38.34.179.78:8448 | ✓ 247ms | ✓ 837ms | ✓ 152ms | ✓ 842ms | ✓ 538ms | http |
| 38.145.218.218:8451 | ✓ 805ms | ✓ 1623ms | ✓ 155ms | ✓ 784ms | ✓ 1113ms | http |
| 38.145.208.169:8452 | ✓ 275ms | 否 | ✓ 855ms | ✓ 710ms | ✓ 933ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1222ms | ✓ 1181ms | ✓ 968ms | http |
| 20.118.221.52:3128 | 否 | ✓ 1376ms | ✓ 232ms | ✓ 1499ms | ✓ 892ms | http |
| 38.34.179.42:8449 | ✓ 280ms | 否 | ✓ 836ms | ✓ 680ms | ✓ 565ms | http |
| 38.145.208.214:8446 | ✓ 217ms | ✓ 1120ms | ✓ 1409ms | ✓ 1295ms | ✓ 510ms | http |
| 103.82.93.219:3128 | ✓ 825ms | 否 | ✓ 969ms | ✓ 1245ms | ✓ 989ms | http |
| 14.225.8.195:3218 | ✓ 1208ms | 否 | ✓ 1023ms | ✓ 1178ms | ✓ 918ms | http |
| 155.117.18.36:25388 | 否 | 否 | ✓ 1043ms | ✓ 1409ms | ✓ 1049ms | http |
| 104.243.46.122:3128 | 否 | 否 | ✓ 1318ms | ✓ 1398ms | ✓ 976ms | http |
| 194.116.191.181:3128 | ✓ 1464ms | ✓ 1851ms | ✓ 1483ms | ✓ 1451ms | ✓ 1380ms | http |
| 45.88.0.114:3128 | ✓ 1542ms | ✓ 1978ms | ✓ 1375ms | 否 | ✓ 1720ms | http |
| 45.88.0.115:3128 | ✓ 1543ms | 否 | ✓ 1352ms | 否 | ✓ 1743ms | http |
| 213.220.62.63:3128 | ✓ 1530ms | 否 | ✓ 1357ms | 否 | ✓ 1726ms | http |
| 45.88.0.99:3128 | ✓ 1536ms | ✓ 1876ms | ✓ 1488ms | 否 | ✓ 1718ms | http |
| 45.88.0.117:3128 | ✓ 1542ms | ✓ 1911ms | ✓ 1436ms | 否 | ✓ 1724ms | http |
| 45.88.0.116:3128 | ✓ 1536ms | 否 | ✓ 1363ms | 否 | ✓ 1711ms | http |
| 213.220.62.62:3128 | ✓ 1535ms | 否 | ✓ 1359ms | 否 | ✓ 1746ms | http |
| 45.88.0.111:3128 | ✓ 1542ms | ✓ 1674ms | ✓ 1678ms | 否 | ✓ 1756ms | http |
| 45.88.0.113:3128 | ✓ 1536ms | ✓ 1752ms | ✓ 1606ms | 否 | ✓ 1738ms | http |
| 45.88.0.98:3128 | ✓ 1534ms | 否 | ✓ 1366ms | 否 | ✓ 1718ms | http |
| 103.231.75.209:3128 | ✓ 1495ms | 否 | ✓ 1126ms | 否 | ✓ 1813ms | http |
| 91.238.104.172:2024 | ✓ 1496ms | 否 | ✓ 1739ms | 否 | ✓ 1657ms | http |
| 91.238.105.64:2024 | ✓ 1497ms | 否 | ✓ 1844ms | 否 | ✓ 1738ms | http |
| 163.61.38.128:3128 | ✓ 1908ms | 否 | ✓ 1653ms | 否 | ✓ 1678ms | http |
| 89.208.106.138:10808 | ✓ 1903ms | ✓ 1695ms | ✓ 1765ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1279ms | ✓ 1890ms | ✓ 1700ms | ✓ 1518ms | 否 | http |
| 137.59.47.73:3128 | ✓ 1106ms | ✓ 1647ms | 否 | ✓ 1484ms | ✓ 1443ms | http |
| 116.80.65.74:3172 | 否 | 否 | ✓ 1455ms | ✓ 1769ms | ✓ 1658ms | http |
| 104.234.0.145:55554 | ✓ 110ms | ✓ 821ms | ✓ 1158ms | ✓ 1833ms | ✓ 610ms | http |
| 170.106.137.214:7890 | ✓ 1041ms | 否 | ✓ 1681ms | ✓ 1297ms | ✓ 1081ms | http |
| 38.244.54.190:31168 | ✓ 1114ms | ✓ 1059ms | ✓ 761ms | 否 | 否 | http |
| 38.34.179.24:8447 | ✓ 749ms | ✓ 1802ms | ✓ 662ms | 否 | 否 | http |
| 38.34.179.99:8446 | ✓ 782ms | ✓ 816ms | ✓ 316ms | ✓ 789ms | ✓ 524ms | http |
| 38.145.208.177:8450 | ✓ 1426ms | ✓ 844ms | ✓ 651ms | ✓ 676ms | 否 | http |
| 38.34.179.74:8447 | ✓ 792ms | ✓ 1510ms | ✓ 589ms | ✓ 708ms | 否 | http |
| 38.145.218.160:8448 | ✓ 770ms | ✓ 1778ms | ✓ 264ms | ✓ 723ms | 否 | http |
| 38.145.218.161:8445 | ✓ 773ms | 否 | ✓ 96ms | ✓ 720ms | 否 | http |
| 45.136.130.170:8446 | ✓ 1057ms | 否 | ✓ 503ms | 否 | ✓ 744ms | http |
| 38.145.220.82:8448 | ✓ 759ms | ✓ 1402ms | ✓ 898ms | ✓ 1897ms | ✓ 1004ms | http |
| 38.145.218.163:8446 | ✓ 1395ms | ✓ 1534ms | ✓ 634ms | 否 | ✓ 1267ms | http |
| 38.145.220.96:8445 | ✓ 1055ms | 否 | ✓ 436ms | 否 | ✓ 1936ms | http |
| 38.34.179.98:8445 | ✓ 920ms | ✓ 1550ms | ✓ 385ms | ✓ 755ms | ✓ 510ms | http |
| 38.145.203.87:8444 | ✓ 683ms | ✓ 667ms | ✓ 104ms | ✓ 704ms | ✓ 536ms | http |
| 38.145.220.72:8445 | ✓ 677ms | ✓ 726ms | ✓ 124ms | ✓ 731ms | ✓ 633ms | http |
| 38.145.218.87:8451 | ✓ 683ms | ✓ 681ms | ✓ 95ms | ✓ 803ms | ✓ 598ms | http |
| 38.34.179.101:8453 | ✓ 683ms | ✓ 735ms | ✓ 79ms | ✓ 699ms | ✓ 591ms | http |
| 38.145.208.172:8446 | ✓ 676ms | ✓ 687ms | ✓ 98ms | ✓ 823ms | ✓ 570ms | http |
| 38.145.203.111:8453 | ✓ 683ms | ✓ 799ms | ✓ 132ms | ✓ 746ms | ✓ 621ms | http |
| 38.34.179.21:8452 | ✓ 682ms | ✓ 628ms | ✓ 155ms | ✓ 702ms | ✓ 570ms | http |
| 38.34.179.31:8453 | ✓ 677ms | ✓ 704ms | ✓ 111ms | ✓ 728ms | ✓ 915ms | http |
| 38.145.208.221:8447 | ✓ 685ms | ✓ 641ms | ✓ 95ms | ✓ 699ms | ✓ 634ms | http |
| 38.145.218.232:8448 | ✓ 679ms | ✓ 795ms | ✓ 117ms | ✓ 846ms | ✓ 676ms | http |
| 38.34.179.88:8446 | ✓ 684ms | ✓ 692ms | ✓ 89ms | ✓ 700ms | ✓ 548ms | http |
| 38.145.220.20:8447 | ✓ 693ms | ✓ 669ms | ✓ 568ms | ✓ 675ms | ✓ 503ms | http |
| 38.145.208.242:8444 | ✓ 694ms | ✓ 1139ms | ✓ 100ms | ✓ 657ms | ✓ 621ms | http |
| 38.34.179.50:8444 | ✓ 695ms | ✓ 653ms | ✓ 371ms | ✓ 709ms | ✓ 541ms | http |
| 38.145.218.212:8448 | ✓ 677ms | ✓ 1263ms | ✓ 81ms | ✓ 669ms | ✓ 524ms | http |
| 38.34.179.49:8445 | ✓ 695ms | ✓ 719ms | ✓ 299ms | ✓ 726ms | ✓ 503ms | http |
| 38.34.179.89:8449 | ✓ 678ms | ✓ 709ms | ✓ 111ms | ✓ 721ms | ✓ 549ms | http |
| 38.34.179.150:8449 | ✓ 680ms | ✓ 1201ms | ✓ 118ms | ✓ 678ms | ✓ 717ms | http |
| 38.34.179.194:8451 | ✓ 696ms | ✓ 1024ms | ✓ 166ms | ✓ 1047ms | ✓ 494ms | http |
| 38.145.203.109:8449 | ✓ 680ms | ✓ 1509ms | ✓ 93ms | ✓ 687ms | ✓ 520ms | http |
| 45.136.130.197:8452 | ✓ 681ms | ✓ 663ms | ✓ 665ms | ✓ 873ms | ✓ 633ms | http |
| 38.145.208.181:8445 | ✓ 694ms | ✓ 661ms | ✓ 599ms | ✓ 1154ms | ✓ 518ms | http |
| 45.136.130.176:8451 | ✓ 724ms | ✓ 1236ms | ✓ 109ms | ✓ 696ms | ✓ 539ms | http |
| 38.145.203.109:8453 | ✓ 684ms | ✓ 1262ms | ✓ 315ms | ✓ 689ms | ✓ 515ms | http |
| 38.34.179.61:8445 | ✓ 678ms | ✓ 1737ms | ✓ 81ms | ✓ 689ms | ✓ 549ms | http |
| 45.136.131.61:8444 | ✓ 683ms | ✓ 777ms | ✓ 124ms | ✓ 1329ms | ✓ 540ms | http |
| 45.136.130.179:8450 | ✓ 681ms | ✓ 1323ms | ✓ 238ms | ✓ 671ms | ✓ 616ms | http |
| 38.34.179.27:8451 | ✓ 1262ms | ✓ 706ms | ✓ 267ms | ✓ 1043ms | ✓ 519ms | http |
| 38.145.208.247:8452 | ✓ 698ms | ✓ 1647ms | ✓ 88ms | ✓ 688ms | ✓ 624ms | http |
| 38.145.208.245:8452 | ✓ 695ms | ✓ 1709ms | ✓ 83ms | ✓ 759ms | ✓ 531ms | http |

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
