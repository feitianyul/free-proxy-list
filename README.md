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

最后更新：2026-03-10 09:43:12 UTC（2026-03-10 17:43:12 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 154.3.236.202:3128 | ✓ 965ms | 否 | ✓ 1259ms | ✓ 1962ms | ✓ 931ms | http |
| 95.3.9.78:3128 | ✓ 1632ms | 否 | 否 | ✓ 1877ms | ✓ 1327ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1789ms | ✓ 1395ms | ✓ 957ms | http |
| 34.101.184.164:3128 | ✓ 1817ms | 否 | ✓ 1629ms | 否 | ✓ 1290ms | http |
| 120.92.212.16:7890 | ✓ 1733ms | ✓ 1332ms | 否 | 否 | ✓ 1803ms | http |
| 190.212.131.238:3128 | ✓ 891ms | ✓ 1662ms | ✓ 701ms | ✓ 1936ms | ✓ 1302ms | http |
| 95.3.9.78:8080 | ✓ 889ms | ✓ 1779ms | ✓ 650ms | ✓ 1645ms | ✓ 1231ms | http |
| 205.209.118.30:3138 | ✓ 769ms | 否 | ✓ 169ms | ✓ 1198ms | ✓ 1269ms | http |
| 91.233.223.147:3128 | ✓ 1035ms | ✓ 1944ms | ✓ 1117ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1003ms | 否 | ✓ 1023ms | ✓ 1484ms | 否 | http |
| 152.42.213.210:8080 | ✓ 1784ms | 否 | ✓ 1704ms | ✓ 1862ms | ✓ 1435ms | http |
| 91.107.141.42:8081 | ✓ 806ms | 否 | ✓ 1644ms | 否 | ✓ 1605ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 625ms | ✓ 1030ms | ✓ 868ms | http |
| 45.136.131.47:8443 | ✓ 904ms | 否 | ✓ 326ms | ✓ 974ms | ✓ 755ms | http |
| 39.104.201.40:7890 | ✓ 1086ms | ✓ 1401ms | ✓ 1100ms | ✓ 1407ms | ✓ 1116ms | http |
| 101.43.255.96:80 | ✓ 1162ms | 否 | ✓ 1067ms | ✓ 1388ms | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1792ms | ✓ 1372ms | ✓ 1395ms | http |
| 190.9.109.198:999 | ✓ 1078ms | ✓ 1428ms | ✓ 1403ms | ✓ 1407ms | ✓ 1305ms | http |
| 194.213.18.200:443 | ✓ 757ms | 否 | ✓ 1993ms | ✓ 1973ms | ✓ 1155ms | http |
| 46.183.25.8:443 | ✓ 1499ms | 否 | 否 | ✓ 1963ms | ✓ 841ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1350ms | ✓ 1332ms | ✓ 1531ms | http |
| 185.191.236.162:3128 | ✓ 1265ms | 否 | ✓ 1755ms | 否 | ✓ 1377ms | http |
| 168.235.110.63:3128 | ✓ 64ms | 否 | ✓ 533ms | ✓ 1030ms | ✓ 1266ms | http |
| 45.136.130.223:8443 | ✓ 813ms | 否 | ✓ 317ms | ✓ 973ms | ✓ 801ms | http |
| 94.176.3.43:7443 | ✓ 1535ms | 否 | ✓ 891ms | ✓ 1718ms | ✓ 1271ms | http |
| 5.101.0.233:3128 | ✓ 644ms | 否 | ✓ 1755ms | 否 | ✓ 1544ms | http |
| 81.70.169.194:80 | ✓ 1053ms | ✓ 1435ms | 否 | ✓ 1409ms | 否 | http |
| 5.252.33.13:2025 | ✓ 1978ms | 否 | ✓ 1168ms | 否 | ✓ 1724ms | http |
| 165.227.5.10:8888 | ✓ 476ms | 否 | ✓ 793ms | ✓ 1012ms | ✓ 1693ms | http |
| 185.115.74.185:8080 | ✓ 1043ms | ✓ 1626ms | ✓ 1463ms | 否 | 否 | http |
| 117.159.239.58:22222 | ✓ 1052ms | ✓ 1252ms | ✓ 1057ms | ✓ 1332ms | ✓ 1020ms | http |
| 121.204.158.249:3128 | ✓ 1042ms | ✓ 1720ms | ✓ 979ms | ✓ 1219ms | ✓ 959ms | http |
| 113.177.131.2:3128 | ✓ 1539ms | 否 | ✓ 1033ms | ✓ 1620ms | ✓ 1132ms | http |
| 159.223.42.219:3128 | ✓ 1770ms | 否 | ✓ 1791ms | ✓ 1541ms | ✓ 1073ms | http |
| 222.228.171.92:8080 | ✓ 1818ms | 否 | ✓ 1500ms | 否 | ✓ 1862ms | http |
| 183.249.5.214:22222 | 否 | ✓ 1223ms | 否 | ✓ 1692ms | ✓ 1025ms | http |
| 136.49.34.18:8888 | ✓ 825ms | 否 | ✓ 1470ms | ✓ 1989ms | ✓ 1504ms | http |
| 86.53.183.16:1080 | ✓ 676ms | 否 | ✓ 1408ms | 否 | ✓ 1875ms | http |
| 162.248.165.72:1080 | ✓ 1072ms | 否 | ✓ 650ms | 否 | ✓ 1656ms | http |
| 103.82.23.118:5171 | ✓ 1791ms | 否 | ✓ 1465ms | 否 | ✓ 1318ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1311ms | ✓ 1443ms | ✓ 1506ms | 否 | http |
| 113.59.32.162:22222 | 否 | ✓ 1510ms | 否 | ✓ 1788ms | ✓ 1205ms | http |
| 178.236.245.17:3128 | ✓ 1288ms | ✓ 1708ms | ✓ 1480ms | 否 | 否 | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1459ms | ✓ 1835ms | ✓ 1325ms | http |
| 117.159.239.49:22222 | ✓ 1148ms | 否 | ✓ 958ms | ✓ 1355ms | ✓ 1058ms | http |
| 120.238.159.234:22222 | ✓ 1145ms | ✓ 1459ms | ✓ 1165ms | ✓ 1408ms | ✓ 1133ms | http |
| 113.59.32.161:22222 | ✓ 1239ms | 否 | ✓ 1245ms | ✓ 1453ms | ✓ 1232ms | http |
| 162.240.154.26:3128 | ✓ 806ms | ✓ 1586ms | ✓ 1637ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 629ms | ✓ 1497ms | ✓ 1730ms | ✓ 1776ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1623ms | ✓ 1050ms | ✓ 1339ms | ✓ 1023ms | http |
| 183.249.5.111:22222 | ✓ 871ms | ✓ 1399ms | 否 | 否 | ✓ 943ms | http |
| 207.254.71.62:8088 | ✓ 554ms | 否 | ✓ 1234ms | ✓ 1751ms | ✓ 1564ms | http |
| 47.77.193.180:1080 | ✓ 839ms | ✓ 1596ms | ✓ 888ms | ✓ 937ms | ✓ 693ms | http |
| 88.80.150.82:8080 | ✓ 1234ms | 否 | 否 | ✓ 1952ms | ✓ 1802ms | https |
| 103.39.51.190:8080 | ✓ 1992ms | 否 | 否 | ✓ 1595ms | ✓ 1815ms | http |
| 120.240.29.173:22222 | ✓ 1391ms | ✓ 1454ms | ✓ 1192ms | ✓ 1456ms | ✓ 1142ms | http |
| 152.42.213.210:80 | 否 | 否 | ✓ 1798ms | ✓ 1375ms | ✓ 1078ms | http |
| 61.155.242.150:5566 | ✓ 988ms | ✓ 1180ms | ✓ 1642ms | ✓ 1876ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1276ms | 否 | 否 | ✓ 1728ms | ✓ 1319ms | http |
| 45.140.147.155:1081 | ✓ 1044ms | ✓ 1373ms | ✓ 1157ms | ✓ 1592ms | ✓ 861ms | http |
| 61.52.131.172:8443 | ✓ 1123ms | ✓ 1333ms | ✓ 1096ms | ✓ 1380ms | ✓ 1065ms | http |
| 59.46.216.131:30001 | ✓ 1102ms | 否 | ✓ 1653ms | ✓ 1434ms | 否 | http |
| 103.35.188.243:3128 | ✓ 241ms | ✓ 1045ms | ✓ 1769ms | ✓ 1157ms | ✓ 867ms | http |
| 178.236.245.59:3128 | ✓ 1295ms | ✓ 1657ms | ✓ 1545ms | ✓ 1927ms | ✓ 1462ms | http |
| 103.3.246.71:3128 | 否 | 否 | ✓ 1139ms | ✓ 1287ms | ✓ 1049ms | http |

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
