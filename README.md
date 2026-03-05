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

最后更新：2026-03-05 14:01:55 UTC（2026-03-05 22:01:55 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 357ms | 否 | ✓ 1800ms | ✓ 1118ms | ✓ 837ms | http |
| 35.225.22.61:80 | ✓ 957ms | 否 | ✓ 986ms | ✓ 958ms | ✓ 892ms | http |
| 20.27.15.49:8561 | ✓ 771ms | ✓ 1415ms | ✓ 839ms | ✓ 1122ms | ✓ 843ms | http |
| 20.210.76.104:8561 | ✓ 814ms | ✓ 1466ms | ✓ 747ms | ✓ 1122ms | ✓ 853ms | http |
| 45.136.198.40:3128 | ✓ 943ms | 否 | ✓ 581ms | ✓ 1423ms | ✓ 1110ms | http |
| 156.225.70.152:39151 | 否 | 否 | ✓ 1772ms | ✓ 1524ms | ✓ 948ms | http |
| 165.227.5.10:8888 | ✓ 451ms | 否 | ✓ 523ms | ✓ 1001ms | ✓ 773ms | http |
| 103.84.95.54:7890 | ✓ 1026ms | 否 | 否 | ✓ 1055ms | ✓ 1018ms | http |
| 188.132.141.249:443 | ✓ 1653ms | ✓ 1893ms | 否 | 否 | ✓ 1780ms | http |
| 20.210.39.153:8561 | ✓ 1468ms | ✓ 1365ms | ✓ 647ms | 否 | 否 | http |
| 20.27.15.111:8561 | ✓ 1455ms | ✓ 1540ms | ✓ 618ms | 否 | 否 | http |
| 20.27.11.248:8561 | ✓ 1460ms | ✓ 1585ms | ✓ 642ms | 否 | 否 | http |
| 20.27.14.220:8561 | ✓ 1489ms | ✓ 1207ms | ✓ 662ms | 否 | 否 | http |
| 168.235.110.63:3128 | ✓ 528ms | 否 | 否 | ✓ 1423ms | ✓ 1038ms | http |
| 91.193.240.157:9877 | ✓ 1205ms | 否 | ✓ 1568ms | ✓ 1855ms | ✓ 1420ms | http |
| 104.129.203.245:10026 | ✓ 321ms | 否 | ✓ 268ms | ✓ 1038ms | ✓ 891ms | http |
| 104.129.203.245:10733 | ✓ 614ms | ✓ 986ms | ✓ 951ms | ✓ 1080ms | ✓ 1366ms | http |
| 104.129.203.245:10139 | ✓ 441ms | ✓ 1620ms | ✓ 1236ms | ✓ 998ms | ✓ 792ms | http |
| 165.225.113.220:11462 | ✓ 1409ms | 否 | ✓ 1004ms | 否 | ✓ 1110ms | http |
| 62.113.119.14:8080 | ✓ 624ms | ✓ 1780ms | ✓ 948ms | 否 | ✓ 1136ms | http |
| 165.225.72.38:10603 | ✓ 598ms | 否 | ✓ 1204ms | ✓ 1536ms | ✓ 1623ms | http |
| 165.225.72.38:10914 | ✓ 1530ms | ✓ 1494ms | ✓ 750ms | 否 | ✓ 1192ms | http |
| 165.225.72.38:11080 | ✓ 413ms | 否 | ✓ 396ms | ✓ 1936ms | ✓ 1031ms | http |
| 20.78.26.206:8561 | ✓ 949ms | ✓ 1611ms | ✓ 647ms | ✓ 1038ms | ✓ 818ms | http |
| 20.78.118.91:8561 | ✓ 952ms | 否 | ✓ 685ms | ✓ 1008ms | ✓ 846ms | http |
| 61.72.221.194:3128 | ✓ 924ms | 否 | ✓ 934ms | 否 | ✓ 1642ms | http |
| 165.225.72.38:10101 | ✓ 500ms | ✓ 1531ms | ✓ 1531ms | ✓ 1442ms | 否 | http |
| 165.225.72.38:10003 | ✓ 437ms | 否 | ✓ 411ms | ✓ 1300ms | ✓ 1025ms | http |
| 165.225.72.38:11405 | ✓ 572ms | ✓ 1567ms | ✓ 793ms | 否 | ✓ 1736ms | http |
| 165.225.72.38:10561 | ✓ 391ms | ✓ 1535ms | ✓ 398ms | ✓ 1818ms | ✓ 1012ms | http |
| 165.225.72.38:10923 | ✓ 399ms | ✓ 1523ms | ✓ 395ms | ✓ 1820ms | ✓ 1045ms | http |
| 165.225.72.38:10869 | ✓ 399ms | ✓ 1501ms | ✓ 416ms | 否 | ✓ 1032ms | http |
| 165.225.72.38:9401 | ✓ 394ms | 否 | ✓ 465ms | ✓ 1284ms | ✓ 1792ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1465ms | 否 | ✓ 1370ms | ✓ 1428ms | http |
| 101.43.255.96:80 | ✓ 1180ms | 否 | ✓ 1580ms | ✓ 1841ms | 否 | http |
| 159.89.31.62:8080 | ✓ 486ms | ✓ 1854ms | ✓ 1946ms | ✓ 1905ms | ✓ 1325ms | http |
| 106.14.205.114:483 | ✓ 1581ms | 否 | 否 | ✓ 1316ms | ✓ 1931ms | http |
| 207.254.71.62:8088 | ✓ 1021ms | 否 | ✓ 1246ms | ✓ 1727ms | ✓ 1512ms | http |
| 120.92.212.16:8890 | ✓ 1185ms | 否 | ✓ 1194ms | 否 | ✓ 1166ms | http |
| 185.191.236.162:3128 | 否 | 否 | ✓ 1206ms | ✓ 1963ms | ✓ 1356ms | http |
| 183.237.195.130:3128 | 否 | 否 | ✓ 1590ms | ✓ 1524ms | ✓ 1300ms | http |
| 120.92.212.16:7890 | ✓ 1508ms | ✓ 1500ms | 否 | 否 | ✓ 1154ms | http |
| 58.220.95.12:11743 | ✓ 1559ms | 否 | 否 | ✓ 1421ms | ✓ 1768ms | http |
| 121.128.121.54:3128 | ✓ 797ms | 否 | 否 | ✓ 1212ms | ✓ 972ms | http |
| 5.75.196.26:40000 | ✓ 845ms | ✓ 1942ms | ✓ 812ms | 否 | 否 | http |
| 58.220.95.11:11023 | ✓ 1110ms | 否 | 否 | ✓ 1755ms | ✓ 1116ms | http |
| 210.223.44.230:3128 | ✓ 783ms | 否 | ✓ 1191ms | ✓ 1141ms | ✓ 869ms | http |
| 120.55.163.237:10086 | ✓ 1081ms | ✓ 1295ms | ✓ 1158ms | ✓ 1424ms | ✓ 1079ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1654ms | ✓ 1924ms | ✓ 1167ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1657ms | ✓ 1901ms | ✓ 1544ms | http |
| 90.84.188.97:8000 | ✓ 943ms | 否 | ✓ 664ms | 否 | ✓ 1809ms | http |
| 192.71.213.85:9091 | ✓ 776ms | 否 | ✓ 1540ms | ✓ 1976ms | 否 | http |
| 165.225.113.220:11233 | 否 | ✓ 1685ms | 否 | ✓ 1525ms | ✓ 1284ms | http |
| 185.243.218.43:49153 | ✓ 841ms | 否 | ✓ 1379ms | ✓ 1745ms | ✓ 1448ms | http |
| 103.39.51.190:8080 | ✓ 1990ms | 否 | 否 | ✓ 1569ms | ✓ 1621ms | http |
| 172.212.68.37:3128 | ✓ 178ms | 否 | ✓ 686ms | ✓ 1252ms | ✓ 1027ms | http |
| 121.237.181.137:8888 | 否 | 否 | ✓ 1069ms | ✓ 1398ms | ✓ 1139ms | http |
| 14.56.177.44:3128 | 否 | 否 | ✓ 1580ms | ✓ 1346ms | ✓ 985ms | http |
| 95.85.252.153:21064 | ✓ 412ms | ✓ 1638ms | ✓ 1332ms | 否 | 否 | http |
| 221.122.91.36:11273 | ✓ 1184ms | ✓ 1403ms | ✓ 1104ms | ✓ 1497ms | ✓ 1177ms | http |
| 88.80.150.82:8080 | ✓ 699ms | ✓ 1740ms | ✓ 788ms | 否 | 否 | https |
| 45.129.141.143:3128 | ✓ 1650ms | ✓ 1888ms | 否 | ✓ 1931ms | ✓ 1487ms | http |
| 81.70.169.194:80 | 否 | ✓ 1427ms | 否 | ✓ 1539ms | ✓ 1527ms | http |
| 103.149.99.128:3128 | 否 | ✓ 1935ms | ✓ 1277ms | ✓ 1459ms | ✓ 1197ms | http |

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
