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

最后更新：2026-04-25 06:06:15 UTC（2026-04-25 14:06:15 UTC+8）

**代理总数：52**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 836ms | ✓ 1058ms | ✓ 890ms | ✓ 929ms | ✓ 712ms | http |
| 113.160.132.26:8080 | ✓ 1564ms | ✓ 1684ms | ✓ 1134ms | ✓ 1248ms | ✓ 1008ms | http |
| 206.206.126.177:2412 | ✓ 1405ms | 否 | ✓ 1968ms | ✓ 1050ms | ✓ 877ms | http |
| 212.58.132.5:8888 | ✓ 1403ms | 否 | ✓ 1405ms | ✓ 1485ms | ✓ 1254ms | http |
| 35.225.22.61:80 | ✓ 666ms | ✓ 1792ms | ✓ 1076ms | ✓ 1184ms | 否 | http |
| 20.127.128.70:8080 | ✓ 855ms | 否 | ✓ 1017ms | ✓ 1671ms | ✓ 1284ms | http |
| 177.93.132.244:3128 | ✓ 1147ms | 否 | ✓ 821ms | 否 | ✓ 1769ms | http |
| 34.71.229.255:3128 | ✓ 767ms | ✓ 1299ms | ✓ 1167ms | ✓ 1032ms | ✓ 772ms | http |
| 8.219.195.129:1080 | ✓ 859ms | 否 | ✓ 902ms | ✓ 1090ms | ✓ 879ms | http |
| 34.96.238.40:8080 | ✓ 1153ms | ✓ 1417ms | ✓ 1247ms | 否 | ✓ 1420ms | http |
| 2.27.54.161:1080 | ✓ 1211ms | 否 | ✓ 1448ms | 否 | ✓ 1566ms | http |
| 94.131.122.129:1081 | ✓ 1203ms | ✓ 1556ms | ✓ 1785ms | 否 | 否 | http |
| 47.85.51.197:1080 | ✓ 209ms | 否 | ✓ 220ms | ✓ 1205ms | ✓ 798ms | http |
| 130.61.174.200:1080 | ✓ 1205ms | 否 | 否 | ✓ 1961ms | ✓ 1526ms | http |
| 208.87.243.199:7878 | ✓ 1422ms | 否 | ✓ 801ms | ✓ 1509ms | 否 | http |
| 45.140.147.82:1082 | 否 | 否 | ✓ 1625ms | ✓ 1940ms | ✓ 1547ms | http |
| 45.140.147.82:1081 | 否 | 否 | ✓ 1605ms | ✓ 1999ms | ✓ 1491ms | http |
| 62.113.119.14:8080 | ✓ 1306ms | 否 | ✓ 691ms | ✓ 1638ms | ✓ 1211ms | http |
| 168.144.75.9:3128 | ✓ 1236ms | 否 | ✓ 1631ms | ✓ 1959ms | ✓ 1549ms | http |
| 161.35.181.96:999 | ✓ 472ms | ✓ 1180ms | ✓ 1295ms | ✓ 1125ms | ✓ 976ms | http |
| 43.132.188.134:443 | ✓ 676ms | ✓ 1015ms | 否 | ✓ 1884ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1235ms | 否 | ✓ 1886ms | ✓ 1714ms | 否 | http |
| 31.207.47.254:3128 | ✓ 1695ms | ✓ 1828ms | ✓ 1912ms | 否 | 否 | http |
| 52.186.152.254:443 | ✓ 669ms | ✓ 1712ms | ✓ 1626ms | 否 | 否 | http |
| 121.230.8.55:1080 | ✓ 1947ms | 否 | 否 | ✓ 1426ms | ✓ 1150ms | http |
| 52.16.215.4:57294 | ✓ 1395ms | 否 | ✓ 1468ms | 否 | ✓ 1780ms | http |
| 34.246.183.20:9023 | ✓ 1848ms | 否 | ✓ 1677ms | ✓ 1708ms | 否 | http |
| 104.248.243.244:3128 | ✓ 581ms | ✓ 1782ms | ✓ 575ms | ✓ 1429ms | ✓ 1077ms | http |
| 103.187.146.151:3128 | ✓ 1173ms | 否 | ✓ 771ms | ✓ 1198ms | ✓ 927ms | http |
| 47.121.114.42:3129 | ✓ 1845ms | ✓ 1624ms | ✓ 1945ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1348ms | 否 | ✓ 896ms | 否 | ✓ 1622ms | http |
| 166.88.61.54:8000 | ✓ 1069ms | ✓ 1280ms | ✓ 1325ms | ✓ 945ms | ✓ 753ms | http |
| 47.84.76.30:1080 | ✓ 864ms | 否 | ✓ 743ms | ✓ 1098ms | ✓ 863ms | http |
| 86.104.74.110:1081 | 否 | 否 | ✓ 1190ms | ✓ 1952ms | ✓ 1311ms | http |
| 34.246.223.187:47434 | ✓ 839ms | 否 | 否 | ✓ 1799ms | ✓ 1572ms | http |
| 202.40.185.146:2025 | 否 | 否 | ✓ 1583ms | ✓ 1924ms | ✓ 1828ms | http |
| 59.46.216.131:30001 | ✓ 1046ms | ✓ 1331ms | 否 | ✓ 1406ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1897ms | 否 | ✓ 1287ms | ✓ 1288ms | ✓ 953ms | http |
| 171.234.134.26:6616 | 否 | 否 | ✓ 1522ms | ✓ 1453ms | ✓ 1785ms | http |
| 162.240.154.26:3128 | ✓ 1447ms | 否 | ✓ 941ms | ✓ 1544ms | ✓ 1614ms | http |
| 82.114.228.67:1080 | 否 | 否 | ✓ 783ms | ✓ 1647ms | ✓ 1263ms | http |
| 103.69.84.106:3131 | ✓ 1782ms | 否 | ✓ 1402ms | ✓ 1180ms | ✓ 935ms | http |
| 47.101.159.19:8899 | 否 | ✓ 1050ms | ✓ 1006ms | ✓ 1263ms | ✓ 915ms | http |
| 47.84.59.16:1080 | ✓ 1762ms | ✓ 1711ms | 否 | ✓ 1106ms | ✓ 848ms | http |
| 42.200.76.16:3888 | ✓ 1046ms | 否 | ✓ 776ms | ✓ 983ms | ✓ 1530ms | http |
| 217.77.102.18:3128 | ✓ 1440ms | 否 | ✓ 1793ms | 否 | ✓ 1529ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1207ms | ✓ 1503ms | ✓ 1388ms | http |
| 52.210.57.38:39455 | ✓ 1152ms | 否 | ✓ 722ms | 否 | ✓ 1569ms | http |
| 18.216.7.129:46163 | ✓ 1492ms | 否 | ✓ 1299ms | 否 | ✓ 1776ms | http |
| 103.39.51.207:8080 | ✓ 1466ms | 否 | 否 | ✓ 1777ms | ✓ 1649ms | http |
| 103.55.225.34:8080 | ✓ 1509ms | 否 | ✓ 1535ms | 否 | ✓ 1398ms | http |
| 130.61.139.145:3128 | ✓ 1664ms | ✓ 1995ms | 否 | ✓ 1938ms | ✓ 1406ms | http |

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
