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

最后更新：2026-05-08 14:49:20 UTC（2026-05-08 22:49:20 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 838ms | ✓ 997ms | ✓ 833ms | ✓ 1119ms | ✓ 1111ms | http |
| 103.147.152.12:1080 | ✓ 1303ms | ✓ 1793ms | ✓ 868ms | ✓ 1689ms | ✓ 1314ms | http |
| 212.224.88.212:443 | ✓ 679ms | 否 | ✓ 1155ms | 否 | ✓ 1549ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1781ms | ✓ 1192ms | ✓ 1299ms | ✓ 1046ms | http |
| 79.137.205.44:40000 | ✓ 848ms | 否 | ✓ 1326ms | 否 | ✓ 1664ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 681ms | ✓ 781ms | ✓ 619ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 681ms | ✓ 785ms | ✓ 622ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 683ms | ✓ 783ms | ✓ 622ms | http |
| 107.150.41.226:18080 | ✓ 341ms | ✓ 1126ms | ✓ 481ms | ✓ 1395ms | ✓ 1293ms | http |
| 45.125.67.37:8443 | ✓ 1496ms | 否 | ✓ 839ms | ✓ 1081ms | ✓ 1059ms | http |
| 168.110.52.228:3128 | ✓ 757ms | 否 | 否 | ✓ 1665ms | ✓ 915ms | http |
| 65.109.125.111:8443 | ✓ 918ms | ✓ 1761ms | ✓ 1070ms | 否 | 否 | http |
| 43.133.44.89:8888 | ✓ 1402ms | ✓ 1680ms | 否 | ✓ 1603ms | ✓ 776ms | http |
| 185.221.237.57:443 | ✓ 1806ms | 否 | ✓ 1662ms | 否 | ✓ 1676ms | http |
| 185.221.237.57:8443 | ✓ 1816ms | ✓ 1762ms | ✓ 1897ms | 否 | ✓ 1715ms | http |
| 107.174.64.143:1080 | ✓ 1700ms | 否 | 否 | ✓ 1768ms | ✓ 1825ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1604ms | ✓ 1542ms | ✓ 1404ms | http |
| 217.76.245.80:999 | ✓ 1156ms | 否 | ✓ 1564ms | ✓ 1487ms | ✓ 1545ms | http |
| 144.34.228.13:3128 | ✓ 104ms | ✓ 579ms | ✓ 246ms | ✓ 1554ms | ✓ 603ms | http |
| 20.27.14.220:8561 | ✓ 1169ms | ✓ 800ms | ✓ 455ms | ✓ 771ms | ✓ 578ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1756ms | ✓ 923ms | ✓ 1500ms | ✓ 1253ms | http |
| 152.32.132.190:7890 | ✓ 1304ms | ✓ 1521ms | 否 | ✓ 1699ms | 否 | http |
| 181.119.97.24:999 | 否 | 否 | ✓ 1099ms | ✓ 1850ms | ✓ 1543ms | http |
| 158.160.215.167:8124 | ✓ 1853ms | 否 | ✓ 1042ms | 否 | ✓ 1755ms | http |
| 185.125.100.115:40000 | ✓ 1371ms | 否 | ✓ 1450ms | ✓ 1976ms | 否 | http |
| 43.156.132.113:3128 | ✓ 919ms | ✓ 1260ms | ✓ 773ms | ✓ 1045ms | ✓ 825ms | http |
| 94.131.118.129:1082 | ✓ 1072ms | ✓ 1295ms | ✓ 1673ms | ✓ 1898ms | 否 | http |
| 64.188.77.26:3128 | ✓ 1739ms | ✓ 1635ms | 否 | 否 | ✓ 1848ms | http |
| 213.111.146.36:18080 | ✓ 1242ms | 否 | ✓ 1519ms | ✓ 1617ms | ✓ 1478ms | http |
| 185.195.71.218:18080 | ✓ 1255ms | 否 | ✓ 1930ms | 否 | ✓ 1802ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1525ms | 否 | ✓ 1252ms | ✓ 1098ms | http |
| 45.153.231.229:8080 | ✓ 881ms | 否 | ✓ 1148ms | 否 | ✓ 1942ms | http |
| 103.157.200.126:3128 | ✓ 1701ms | 否 | ✓ 1437ms | 否 | ✓ 1541ms | http |
| 86.104.72.219:1081 | ✓ 881ms | ✓ 1140ms | ✓ 511ms | 否 | ✓ 1093ms | http |
| 138.197.68.35:4857 | ✓ 880ms | ✓ 1586ms | ✓ 434ms | 否 | 否 | http |
| 220.121.143.33:3128 | ✓ 1560ms | ✓ 1608ms | ✓ 709ms | ✓ 958ms | ✓ 1043ms | http |
| 64.188.77.221:3128 | ✓ 1483ms | ✓ 1654ms | ✓ 1102ms | 否 | 否 | http |
| 91.242.229.129:8092 | ✓ 1535ms | 否 | 否 | ✓ 1794ms | ✓ 1522ms | http |
| 94.131.118.129:1081 | ✓ 659ms | 否 | ✓ 573ms | 否 | ✓ 1424ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1990ms | ✓ 1189ms | ✓ 1853ms | http |
| 84.201.138.232:3128 | ✓ 1528ms | 否 | ✓ 1145ms | 否 | ✓ 1582ms | http |
| 173.212.246.157:3128 | ✓ 1002ms | ✓ 1788ms | ✓ 1150ms | 否 | 否 | http |
| 116.171.106.111:3443 | 否 | ✓ 1829ms | 否 | ✓ 1670ms | ✓ 1902ms | http |
| 45.13.116.210:21537 | ✓ 550ms | ✓ 1513ms | ✓ 528ms | ✓ 1472ms | ✓ 1091ms | http |
| 86.104.72.220:1081 | ✓ 418ms | ✓ 1116ms | ✓ 1177ms | ✓ 1248ms | ✓ 892ms | http |
| 86.104.72.220:1082 | ✓ 415ms | ✓ 1689ms | ✓ 588ms | ✓ 1264ms | ✓ 1927ms | http |
| 91.233.223.147:3128 | ✓ 920ms | 否 | ✓ 1069ms | 否 | ✓ 1700ms | http |
| 103.182.52.15:8080 | ✓ 1914ms | 否 | 否 | ✓ 1595ms | ✓ 1912ms | http |
| 101.32.244.83:8080 | 否 | 否 | ✓ 946ms | ✓ 1265ms | ✓ 1263ms | http |
| 121.43.196.210:8222 | ✓ 893ms | ✓ 1042ms | ✓ 894ms | ✓ 1127ms | ✓ 884ms | http |
| 121.43.196.213:8222 | ✓ 907ms | ✓ 1053ms | ✓ 903ms | ✓ 1226ms | ✓ 892ms | http |
| 86.104.72.219:1082 | ✓ 1268ms | 否 | ✓ 254ms | ✓ 1300ms | ✓ 946ms | http |
| 157.0.142.246:10057 | 否 | 否 | ✓ 1607ms | ✓ 1354ms | ✓ 1117ms | http |
| 20.210.39.153:8561 | ✓ 1558ms | 否 | ✓ 1211ms | ✓ 1455ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1557ms | ✓ 1961ms | ✓ 1250ms | ✓ 1466ms | 否 | http |
| 101.32.243.189:80 | ✓ 1341ms | 否 | ✓ 1566ms | 否 | ✓ 1245ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1631ms | ✓ 1112ms | ✓ 1693ms | ✓ 1304ms | http |
| 84.47.150.125:1080 | ✓ 1132ms | 否 | 否 | ✓ 1894ms | ✓ 1510ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1810ms | ✓ 1789ms | ✓ 1978ms | http |
| 59.46.216.131:30001 | ✓ 1003ms | 否 | 否 | ✓ 1414ms | ✓ 1878ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1223ms | ✓ 1014ms | 否 | ✓ 1729ms | http |
| 117.122.240.82:3338 | 否 | ✓ 1643ms | ✓ 1136ms | 否 | ✓ 1140ms | http |
| 3.101.133.120:80 | ✓ 339ms | 否 | ✓ 1380ms | ✓ 1155ms | ✓ 1001ms | http |
| 49.65.127.215:3128 | ✓ 1958ms | 否 | ✓ 1021ms | ✓ 1087ms | ✓ 889ms | http |
| 152.70.91.193:40000 | ✓ 1545ms | 否 | 否 | ✓ 1275ms | ✓ 1424ms | http |
| 64.181.254.251:10443 | ✓ 125ms | ✓ 739ms | ✓ 94ms | ✓ 740ms | ✓ 596ms | http |
| 180.93.35.207:7777 | ✓ 1567ms | 否 | ✓ 1569ms | ✓ 1662ms | 否 | http |
| 147.45.178.211:14658 | ✓ 1272ms | 否 | ✓ 1648ms | 否 | ✓ 1976ms | http |
| 223.16.170.103:80 | ✓ 1080ms | 否 | ✓ 839ms | ✓ 1053ms | ✓ 1099ms | http |
| 103.35.190.69:1081 | ✓ 424ms | ✓ 1770ms | ✓ 285ms | ✓ 1293ms | ✓ 906ms | http |
| 223.16.170.103:3128 | ✓ 1063ms | 否 | ✓ 851ms | ✓ 1043ms | 否 | http |
| 45.59.122.132:80 | 否 | 否 | ✓ 1444ms | ✓ 1960ms | ✓ 1171ms | http |
| 62.113.119.14:8080 | ✓ 763ms | 否 | ✓ 745ms | ✓ 1670ms | ✓ 1177ms | http |
| 45.134.39.140:3333 | ✓ 1483ms | 否 | ✓ 925ms | ✓ 1897ms | 否 | http |
| 178.63.155.151:8888 | ✓ 1132ms | 否 | ✓ 1179ms | ✓ 1707ms | ✓ 1436ms | http |
| 158.160.215.167:8127 | 否 | 否 | ✓ 1062ms | ✓ 1929ms | ✓ 1753ms | http |
| 61.52.131.172:8443 | ✓ 883ms | ✓ 1147ms | ✓ 903ms | 否 | ✓ 1467ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1498ms | ✓ 1024ms | ✓ 1275ms | 否 | http |

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
